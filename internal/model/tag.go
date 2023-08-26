package model

import (
	"blog-service/pkg/app"

	"github.com/jinzhu/gorm"
)

type TagSwagger struct {
	List  []*Tag
	Pager *app.Pager
}

// 標籤
type Tag struct {
	*Model        //引用公共欄位
	Name   string `json:"name"`
	State  uint8  `json:"state"`
}

func (t Tag) TableName() string {
	return "blog_tag"
}

func (t Tag) Count(db *gorm.DB) (int, error) {
	var count int
	//如果name欄位不是空白，添加name欄位過濾
	if t.Name != "" {
		db = db.Where("name = ?", t.Name)
	}
	// 添加state欄位過濾
	db = db.Where("state = ?", t.State)
	//使用Tag資料模型，並用is_del為0過濾並計算數量存至count變數，最後確認是否返回錯誤
	err := db.Model(&t).Where("is_del = ?", 0).Count(&count).Error
	if err != nil {
		return 0, err
	}

	return count, nil
}

func (t Tag) List(db *gorm.DB, pageOffset, pageSize int) ([]*Tag, error) {
	var tags []*Tag
	var err error
	if pageOffset >= 0 && pageSize > 0 {
		db = db.Offset(pageOffset).Limit(pageSize)
	}
	if t.Name != "" {
		db.Where("name= ?", t.Name)
	}
	db = db.Where("state = ?", t.State)
	if err = db.Where("is_del = ?", 0).Find(&tags).Error; err != nil {
		return nil, err
	}
	return tags, nil
}

func (t Tag) Create(db *gorm.DB) error {
	return db.Create(&t).Error
}

func (t Tag) Update(db *gorm.DB, values interface{}) error {
	err := db.Model(t).Where("id = ? AND is_del = ?", t.ID, 0).Updates(values).Error
	if err != nil {
		return err
	}

	return nil
}

func (t Tag) Delete(db *gorm.DB) error {
	return db.Where("id = ? AND is_del =?", t.Model.ID, 0).Delete(&t).Error
}
