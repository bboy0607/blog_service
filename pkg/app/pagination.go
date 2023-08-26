package app

import (
	"blog-service/global"
	"blog-service/pkg/convert"

	"github.com/gin-gonic/gin"
)

// 從query中取page值，如果page<=0則返回1
func GetPage(c *gin.Context) int {
	page := convert.StrTo(c.Query("page")).MustInt()
	if page <= 0 {
		return 1
	}
	return page
}

// 從query取page_size值，如果pagesize<=0則返回DefaultPageSize，如果大於MaxPageSize，則返回MaxPageSize
func GetPageSize(c *gin.Context) int {
	pageSize := convert.StrTo(c.Query("page_size")).MustInt()
	if pageSize <= 0 {
		return global.AppSetting.DefaultPageSize
	}
	if pageSize > global.AppSetting.MaxPageSize {
		return global.AppSetting.MaxPageSize
	}
	return pageSize
}

func GetPageOffset(page, pageSize int) int {
	result := 0
	if page > 0 {
		result = (page - 1) * pageSize
	}
	return result
}
