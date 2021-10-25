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
		rArticle.DELETE("/:id", article.DeleteArticle)
	}

	r.Run(utils.HttpPort)
}
