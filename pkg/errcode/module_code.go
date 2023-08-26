package errcode

var (
	ErrorGetTagListFail = NewErrorCode(20010001, "取得標籤列表失敗")
	ErrorCreateTagFail  = NewErrorCode(20010002, "建立標籤失敗")
	ErrorUpdateTagFail  = NewErrorCode(20010003, "更新標籤失敗")
	ErrorDeleteTagFail  = NewErrorCode(20010004, "刪除標籤失敗")
	ErrorCountTagFail   = NewErrorCode(20010005, "統計標籤失敗")
	ErrorUploadFileFail = NewErrorCode(20030001, "上傳檔案失敗")
)
