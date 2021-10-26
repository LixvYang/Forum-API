package routes

import (
	// "net/http"
	"mixindev/api/v1/article"
	"mixindev/api/v1/user"
	"mixindev/utils"

	"github.com/gin-gonic/gin"
)

func InitRouter() {
	gin.SetMode(utils.AppMode)
	r := gin.Default()

	rUser := r.Group("v1/user") 
	{
		rUser.POST("",user.AddUser)
		rUser.GET("/:id",user.GetUserById)
		rUser.GET("",user.ListUsers)
		rUser.DELETE("/:id",user.DeleteUser)
		rUser.PUT("/:id",user.UpdateUser)
	}

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
