package setting

import "github.com/spf13/viper"

type Setting struct {
	vp *viper.Viper
}

//建立一個含有已讀入設定檔vp的新Setting結構體函數
func NewSetting() (*Setting, error) {
	vp := viper.New()            //初始化一個新viper.Viper的指針
	vp.SetConfigName("config")   //設定設定檔的名稱
	vp.AddConfigPath("configs/") //設定設定檔的存放路徑
	vp.SetConfigType("yaml")     //設定設定檔的檔案類型
	err := vp.ReadInConfig()     //讀入設定檔，並返回一個錯誤
	if err != nil {
		return nil, err
	}

	return &Setting{vp: vp}, err
}
