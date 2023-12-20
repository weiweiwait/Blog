package routers

import (
	"blog-service/docs"
	v1 "blog-service/internal/routers/api/v1"
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func NewRouter() *gin.Engine {
	r := gin.New()
	r.Use(gin.Logger())
	r.Use(gin.Recovery())
	docs.SwaggerInfo.BasePath = ""
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	article := v1.NewArticle()
	tag := v1.NewTag()
	apiv1 := r.Group("/api/v1")
	{
		//新增标签
		apiv1.POST("/tags", tag.Create)
		//删除指定标签
		apiv1.DELETE("/tags/:id", tag.Delete)
		//更新指定标签
		apiv1.PUT("/tags/:id", tag.Update)
		apiv1.PATCH("/tags/:id/state", tag.Update)
		//获取标签列表
		apiv1.GET("/tags", tag.List)

		//新增文章
		apiv1.POST("/articles", article.Create)
		//删除指定文章
		apiv1.DELETE("/articles/:id", article.Delete)
		//更新指定文章
		apiv1.PUT("/articles/:id", article.Update)
		apiv1.PATCH("/articles/:id/state", article.Update)
		//获取指定文章
		apiv1.GET("/articles/:id", article.Get)
		//获取文章列表
		apiv1.GET("/articles", article.List)
	}

	return r
}
