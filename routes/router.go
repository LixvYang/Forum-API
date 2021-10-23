package routes

import (
	// "net/http"
	"mixindev/utils"
	"mixindev/api/v1/article"
	"github.com/gin-gonic/gin"
)

func InitRouter()  {
	gin.SetMode(utils.AppMode)
	r := gin.Default()

	//文章相关
	rArticle := r.Group("v1/article")
	{
		rArticle.GET("/:id",article.GetArticleById)
		rArticle.GET("/ping", func(c *gin.Context) {
			c.JSON(200, gin.H{
				"message": "pong",
			})
		})
	}
	

	r.Run(utils.HttpPort)
}