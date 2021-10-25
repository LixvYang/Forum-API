package routes

import (
	// "net/http"
	"mixindev/api/v1/article"
	"mixindev/utils"

	"github.com/gin-gonic/gin"
)

func InitRouter() {
	gin.SetMode(utils.AppMode)
	r := gin.Default()

	//文章相关
	rArticle := r.Group("v1/article")
	{
		rArticle.GET("/:id", article.GetArticleById)
		rArticle.GET("",article.GetArticlesList)
		rArticle.DELETE("/:id", article.DeleteArticle)
		rArticle.POST("/add",article.AddArticle)
	}

	r.Run(utils.HttpPort)
}
