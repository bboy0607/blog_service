package routers

import (
	"blog-service/global"
	"blog-service/internal/middleware"
	"blog-service/internal/routers/api"
	v1 "blog-service/internal/routers/api/v1"
	"blog-service/pkg/limiter"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"

	_ "blog-service/docs"

	ginSwagger "github.com/swaggo/gin-swagger"
	"github.com/swaggo/gin-swagger/swaggerFiles"
)

var methodLimiters = limiter.NewMMethodLimiter().AddBuckets(
	limiter.LimiterBucketRule{
		Key:          "/auth",
		FillInterval: time.Second,
		Capacity:     10,
		Quantum:      10,
	},
)

func NewRouter() *gin.Engine {
	r := gin.New()
	if global.ServerSetting.RunMode == "debug" {
		r.Use(gin.Logger())   //使用Logger中間層
		r.Use(gin.Recovery()) //使用Recovery中間層
	} else {
		r.Use(middleware.AccessLog())
		r.Use(middleware.Recovery())
	}

	r.Use(middleware.RateLimiter(methodLimiters))
	r.Use(middleware.ContextTimeout(global.AppSetting.DefaultContextTimeout))
	r.Use(middleware.Translations()) //翻譯中介層
	r.Use(middleware.Tracing())
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler)) //swagger文件
	r.GET("/auth", api.GetAuth)

	tag := v1.NewTag()
	article := v1.NewArticle()
	upload := api.NewUpload()
	r.POST("/upload/file", upload.UploadFile)
	r.StaticFS("/static", http.Dir(global.AppSetting.UploadSavePath)) //http.Dir創建一個FileSystem

	apiv1 := r.Group("/api/v1")
	apiv1.Use(middleware.JWT())
	{
		//新增標籤
		apiv1.POST("/tags", tag.Create)
		//取得標籤列表
		apiv1.GET("/tags", tag.List)
		//更新指定標籤
		apiv1.PUT("/tags/:id", tag.Update)
		//更新指定標籤，某個欄位
		apiv1.PATCH("/tags/:id", tag.Update)
		//刪除指定標籤
		apiv1.DELETE("/tags/:id", tag.Delete)

		//新增文章
		apiv1.POST("/articles", article.Create)
		//取得文章列表
		apiv1.GET("/articles", article.List)
		//取得指定文章
		apiv1.GET("/articles/:id", article.Get)
		//更新指定文章
		apiv1.PUT("/articles/:id", article.Update)
		//更新指定文章欄位
		apiv1.PATCH("/articles/:id", article.Update)
		//刪除指定文章
		apiv1.DELETE("/articles/:id", article.Delete)
	}
	return r
}
