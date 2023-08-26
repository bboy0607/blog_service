package v1

import "github.com/gin-gonic/gin"

type Article struct{}

func NewArticle() Article {
	return Article{}
}

// 取得文章的路由方法
func (a Article) Get(c *gin.Context) {}

// 取得文章清單的路由方法
func (a Article) List(c *gin.Context) {}

// 創建文章的路由方法
func (a Article) Create(c *gin.Context) {}

// 更新文章的路由方法
func (a Article) Update(c *gin.Context) {}

// 刪除文章的路由方法
func (a Article) Delete(c *gin.Context) {}
