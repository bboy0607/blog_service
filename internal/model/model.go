package model

import (
	"blog-service/global"
	"blog-service/pkg/setting"
	"fmt"
	"time"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
)

// 公共欄位
type Model struct {
	ID         uint32 `gorm:"primary_key" json:"id"`
	CreatedOn  uint32 `json:"created_on"`
	CreatedBy  string `json:"created_by"`
	ModifiedOn uint32 `json:"modified_on"`
	ModifiedBy string `json:"modified_by"`
	DeletedOn  uint32 `json:"deleted_on"`
	IsDel      uint8  `json:"is_del"`
}

func NewDBEngine(databaseSetting *setting.DatabaseSettingS) (*gorm.DB, error) {
	//讀取參數結構體中的設定值
	s := fmt.Sprintf("%v:%v@tcp(%v)/%v?charset=%v&parseTime=%t&loc=Local",
		databaseSetting.Username,
		databaseSetting.Password,
		databaseSetting.Host,
		databaseSetting.DBName,
		databaseSetting.Charset,
		databaseSetting.ParseTime,
	)
	//依設定值init一個新的DB連線
	db, err := gorm.Open(databaseSetting.DBType, s)
	if err != nil {
		return nil, err
	}

	//如果Server設定區段的RunMode是debug，則啟用DB的詳細Log
	if global.ServerSetting.RunMode == "debug" {
		db.LogMode(true)
	}

	//啟用單數TableName
	db.SingularTable(true)

	//註冊回呼行為
	db.Callback().Create().Replace("gorm:update_time_stamp", updateTimeStampForCreateCallback)
	db.Callback().Update().Replace("gorm:update_time_stamp", updateTimeStampForUpdateCallback)
	db.Callback().Delete().Replace("gorm:delete", deleteCallback)

	//設定最大閒置連線數與最大連線數
	db.DB().SetMaxIdleConns(databaseSetting.MaxIdelConns)
	db.DB().SetMaxOpenConns(databaseSetting.MaxOpenConns)

	return db, nil
}

// 創建用回呼
func updateTimeStampForCreateCallback(scope *gorm.Scope) {
	if !scope.HasError() {
		nowTime := time.Now().Unix()
		//這裡尋找 CreatedOn 欄位是否存在於資料模型中
		if createTimeField, ok := scope.FieldByName("CreatedOn"); ok {
			//如果CreatedOn欄位為空，就設定上現在時間戳
			if createTimeField.IsBlank {
				_ = createTimeField.Set(nowTime)
			}
		}

		//這裡尋找 ModifiedOn 欄位是否存在於資料模型中
		if modifyTimeFiled, ok := scope.FieldByName("ModifiedOn"); ok {
			//如果ModifiedOn欄位為空，就設定上現在時間戳
			if modifyTimeFiled.IsBlank {
				_ = modifyTimeFiled.Set(nowTime)
			}
		}
	}
}

// 更新用回呼
func updateTimeStampForUpdateCallback(scope *gorm.Scope) {
	//scope.Get取得目前設定的標誌 gorm:update_column
	if _, ok := scope.Get("gorm:update_column"); !ok {
		//如果不存在，即沒有自訂設定gorm:update_column，則在更新回呼內設定預設欄位 ModifiedOn的值為目前的時間戳記
		_ = scope.SetColumn("ModifiedOn", time.Now().Unix())
	}
}

// 刪除行為的回呼
func deleteCallback(scope *gorm.Scope) {
	if !scope.HasError() {
		var extraOption string
		if str, ok := scope.Get("gorm:delete_option"); ok {
			extraOption = fmt.Sprint(str)
		}

		deletedOnField, hasDeletedOnField := scope.FieldByName("DeletedOn")
		isDelField, hasIsDelField := scope.FieldByName("IsDel")
		if !scope.Search.Unscoped && hasDeletedOnField && hasIsDelField {
			now := time.Now().Unix()
			scope.Raw(fmt.Sprintf(
				"UPDATE %v SET %v=%v,%v=%v%v%v",
				scope.QuotedTableName(),
				scope.Quote(deletedOnField.DBName),
				scope.AddToVars(now),
				scope.Quote(isDelField.DBName),
				scope.AddToVars(1),
				addExtraSpaceIfExist(scope.CombinedConditionSql()),
				addExtraSpaceIfExist(extraOption),
			)).Exec()
		} else {
			scope.Raw(fmt.Sprintf(
				"DELETE FROM %v%v%v",
				scope.QuotedTableName(),
				addExtraSpaceIfExist(scope.CombinedConditionSql()),
				addExtraSpaceIfExist(extraOption),
			)).Exec()
		}
	}
}

func addExtraSpaceIfExist(str string) string {
	if str != "" {
		return " " + str
	}
	return ""
}
