package upload

import (
	"blog-service/global"
	"blog-service/pkg/util"
	"io"
	"io/ioutil"
	"mime/multipart"
	"os"
	"path"
	"strings"
)

type FileType int

const TypeImage FileType = iota + 1

func GetFileName(name string) string {
	ext := GetFileExt(name)
	//移除附檔名的部分，獲取檔名
	fileName := strings.TrimSuffix(name, ext)
	//對檔名進行 MD5 雜湊編碼，獲得編碼後的值
	fileName = util.EncodeMD5(fileName)
	//ex: 5a105e8b9d40e1329780d62ea2265d8a.jpg
	return fileName + ext
}

// 函數 GetFileExt 是用來獲取檔案名稱的副檔名（擴展名）部分
func GetFileExt(name string) string {
	return path.Ext(name)
}

func GetSavePath() string {
	return global.AppSetting.UploadSavePath
}

// 這個函數 CheckSavePath 是用來檢查指定路徑是否存在的函數。具體來說，它會檢查給定的路徑 dst 是否存在。如果該路徑存在，則返回 false，表示路徑已經存在；如果該路徑不存在，則返回 true，表示路徑不存在。
func CheckSavePath(dst string) bool {
	//使用os.Stat確認路徑是否存在，存在則返回Stat，不存在則返回錯誤
	_, err := os.Stat(dst)
	//判斷錯誤是否表示路徑不存在，如果是路徑不存在，則返回True，否則就是False
	return os.IsNotExist(err)
}

// 檢查副檔名是否有在設定檔"UploadImageAllowExts"的允許清單內
func CheckContainExt(t FileType, name string) bool {
	ext := GetFileExt(name)
	ext = strings.ToUpper(ext)
	switch t {
	case TypeImage:
		for _, allowExt := range global.AppSetting.UploadImageAllowExts {
			if strings.ToUpper(allowExt) == ext {
				return true
			}
		}
	}

	return false
}

// 檢查是否超出最大檔案大小，如果是則返回True
func CheckMaxSize(t FileType, f multipart.File) bool {
	//使用 ioutil.ReadAll(f) 讀取 multipart.File 對象中的內容，並將其全部讀取到一個 byte slice 中
	content, _ := ioutil.ReadAll(f)
	//計算讀取到的內容的長度，即檔案的大小
	size := len(content)
	switch t {
	case TypeImage:
		if size >= global.AppSetting.UploadImageMaxSize*1024*1024 {
			return true
		}
	}

	return false
}

// 檢查檔案許可權是否足夠，如果有權限不足的錯誤就返回True
func CheckPermission(dst string) bool {
	//使用 os.Stat(dst) 函數獲取指定路徑的檔案資訊
	_, err := os.Stat(dst)
	//檢查是否有錯誤發生，如果有錯誤（例如權限不足），則使用 os.IsPermission(err) 函數來判斷錯誤是否表示權限問題
	return os.IsPermission(err)
}

func CreateSavePath(dst string, perm os.FileMode) error {
	err := os.MkdirAll(dst, perm)
	if err != nil {
		return err
	}

	return nil
}

// 這個函數的目的是將上傳的檔案複製到指定的目的地路徑，並在過程中處理錯誤，確保檔案儲存操作的完整性和安全性。
func SaveFile(file *multipart.FileHeader, dst string) error {
	src, err := file.Open()
	if err != nil {
		return err
	}
	defer src.Close()

	out, err := os.Create(dst)
	if err != nil {
		return err
	}
	defer out.Close()

	_, err = io.Copy(out, src)
	return err
}
