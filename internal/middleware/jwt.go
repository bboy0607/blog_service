package middleware

import (
	"blog-service/pkg/app"
	"blog-service/pkg/errcode"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
)

func JWT() gin.HandlerFunc {
	return func(c *gin.Context) {
		var (
			token string
			ecode = errcode.Success
		)
		//從query中抓取token的值
		if s, exist := c.GetQuery("token"); exist {
			token = s
			//如果沒有，則抓取Header中token值
		} else {
			token = c.GetHeader("token")
		}
		//如果token為空，返回參數錯誤
		if token == "" {
			ecode = errcode.InvalidParms
			//使用寫好的ParseToken解析token
		} else {
			_, err := app.ParseToken(token)
			if err != nil {
				//根據錯誤類型進行處理
				switch err.(*jwt.ValidationError).Errors {
				case jwt.ValidationErrorExpired:
					ecode = errcode.UnauthorizedTokenTimeout
				default:
					ecode = errcode.UnauthorizedTokenError
				}
			}
		}

		//最後依ecode判斷，如果不是Success則返回錯誤回應
		if ecode != errcode.Success {
			response := app.NewResponse(c)
			response.ToErrorResponse(ecode)
			c.Abort()
			return
		}
		//如果ecode是Success則繼續執行下一個gin.HandlerFunc
		c.Next()
	}
}
