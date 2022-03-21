package routers

import (
	"github.com/gin-gonic/gin"
	ginSwagger "github.com/swaggo/gin-swagger"
	"github.com/swaggo/gin-swagger/swaggerFiles"
	_ "myproject/docs"
	"myproject/global"
	"myproject/internal/middleware"
	"myproject/internal/routers/api"
	"myproject/internal/routers/api/v1"
	"myproject/pkg/limiter"
	"net/http"
	"time"
)

var methodLimiters = limiter.NewMethodLimiter().AddBuckets(limiter.LimiterBucketRule{
	Key:          "/auth",
	FillInterval: time.Second,
	Capacity:     10,
	Quantum:      10,
})

func NewRouter() *gin.Engine{
	r := gin.New()
	if global.ServerSetting.RunMode == "debug"{
		r.Use(gin.Logger())
		r.Use(gin.Recovery())
	}else{
		r.Use(middleware.AccessLog())
		r.Use(middleware.Recovery())
	}
	r.Use(middleware.RateLimiter(methodLimiters))
	r.Use(middleware.ContextTimeout(global.AppSetting.DefaultContextTimeout))
	r.Use(middleware.Translations())
	r.GET("/swagger/*any",ginSwagger.WrapHandler(swaggerFiles.Handler))
	r.GET("/auth",api.GetAuth )
	tag := v1.NewTag()
	article := v1.NewArticle()
	//上传图片服务
	upload := api.NewUpload()
	r.POST("/upload/file",upload.UploadFile)
	//访问静态文件资源服务
	r.StaticFS("/static",http.Dir(global.AppSetting.UploadSavePath))
	//针对各请求需要有对应的Request结构
	apiv1 := r.Group("/api/v1")
	apiv1.Use(middleware.JWT())
	{
		//标签路由
		apiv1.POST("/tags",tag.Create) //新增标签
		apiv1.DELETE("/tags/:id",tag.Delete)
		apiv1.PUT("/tags/:id",tag.Update) //更新指定标签
		apiv1.PATCH("/tags/:id/state",tag.Update)
		apiv1.GET("/tags",tag.List) //获取标签列表

		//文章路由
		apiv1.POST("/articles",article.Create) //新增文章
		apiv1.DELETE("/articles/:id",article.Delete) //更新指定文章
		apiv1.PUT("/articles/:id",article.Update)
		apiv1.PATCH("/articles/:id/state",article.Update)
		apiv1.GET("/articles/:id",article.Get) //获取指定文章
		apiv1.GET("/articles",article.List) //获取文章列表
	}
	return r
}
