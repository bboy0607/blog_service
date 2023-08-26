package errcode

var (
	Success                   = NewErrorCode(0, "成功")
	ServerError               = NewErrorCode(10000000, "服務內部錯誤")
	InvalidParms              = NewErrorCode(10000001, "導入參數錯誤")
	NotFound                  = NewErrorCode(10000002, "找不到")
	UnauthorizedAuthNotExist  = NewErrorCode(10000003, "驗證失敗,找不到對應的AppKey")
	UnauthorizedTokenError    = NewErrorCode(10000004, "驗證失敗,找不到對應的Token錯誤")
	UnauthorizedTokenTimeout  = NewErrorCode(10000005, "驗證失敗,Token逾時")
	UnauthorizedTokenGenerate = NewErrorCode(10000006, "驗證失敗,Token產生失敗")
	TooManyRequests           = NewErrorCode(10000007, "請求過多")
)
