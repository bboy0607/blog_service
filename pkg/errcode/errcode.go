package errcode

import (
	"fmt"
	"net/http"
)

type Error struct {
	code    int
	msg     string
	details []string
}

// 定義一個package level的Map變數來儲存新增的Error Code
var codes = map[int]string{}

func NewErrorCode(code int, msg string) *Error {
	//使用上面定義的codes的map變數，確認是否有重複的code，如果有的話，返回"已經有這個code"的panic
	if _, ok := codes[code]; ok {
		panic(fmt.Sprintf("code: %v已經存在,請更換一個", code))
	}
	//如果沒有，執行賦值的動作
	codes[code] = msg
	return &Error{code: code, msg: msg}
}

// 定義返回Error結構體的字串方法
func (e *Error) Error() string {
	return fmt.Sprintf("錯誤: %v, 錯誤訊息: %v", e.code, e.msg)
}

// 定義返回Error Code的方法
func (e *Error) Code() int {
	return e.code
}

// 定義返回Error Msg的方法
func (e *Error) Msg() string {
	return e.msg
}

//定義返回Error 帶有msg的格式化訊息的方法
/* 範例:
myErr := MyError{msg: "Error occurred: %s - Code: %d"}
formattedMsg := myErr.Msgf([]interface{}{"Something went wrong", 500})
fmt.Println(formattedMsg)
*/
func (e *Error) Msgf(args []interface{}) string {
	return fmt.Sprintf(e.msg, args...)
}

// 返回Error結構體中的detail 字串切片方法
func (e *Error) Details() []string {
	return e.details
}

// 返回一個帶有Detail的新Error結構體方法
func (e *Error) WithDetails(details ...string) *Error {
	newError := *e
	newError.details = []string{}
	for _, d := range details {
		newError.details = append(newError.details, d)
	}
	return &newError
}

// 內部錯誤碼與HTTP狀態碼的轉換方法
func (e *Error) StatusCode() int {
	switch e.Code() {
	case Success.Code():
		return http.StatusOK //對應http狀態 200
	case ServerError.Code():
		return http.StatusInternalServerError //對應http狀態 500
	case InvalidParms.Code():
		return http.StatusBadRequest //對應http狀態 400
	case UnauthorizedAuthNotExist.Code():
		fallthrough //穿透到後面的case，返回http狀態 401 http.StatusUnauthorized
	case UnauthorizedTokenError.Code():
		fallthrough //穿透到後面的case，返回http狀態 401 http.StatusUnauthorized
	case UnauthorizedTokenGenerate.Code():
		fallthrough //穿透到後面的case，返回http狀態 401 http.StatusUnauthorized
	case UnauthorizedTokenTimeout.Code():
		return http.StatusUnauthorized
	case TooManyRequests.Code():
		return http.StatusTooManyRequests //對應http狀態 429
	}
	return http.StatusInternalServerError //如果沒對應到任何一個，返回http狀態500
}
