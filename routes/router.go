package routes

import (
	"mixindev/api/v1/article"
	"mixindev/api/v1/user"
	"mixindev/api/v1/category"
	"mixindev/utils"

	"github.com/gin-gonic/gin"
)

func InitRouter() {
	gin.SetMode(utils.AppMode)
	r := gin.Default()

	r.POST("v1/login", user.Login)

	rUser := r.Group("v1/user")
	{
		rUser.POST("", user.AddUser)
		rUser.GET("/:id", user.GetUserById)
		rUser.GET("", user.ListUsers)
		rUser.DELETE("/:id", user.DeleteUser)
		rUser.PUT("/:id", user.UpdateUser)
	}

	//文章相关
	rArticle := r.Group("v1/article")
	{
		rArticle.GET("/:id", article.GetArticleById)
		rArticle.GET("", article.GetArticlesList)
		rArticle.DELETE("/:id", article.DeleteArticle)
		rArticle.POST("/add", article.AddArticle)
	}

	rCategory := r.Group("v1/category")
	{
		rCategory.POST("", category.AddCategory)
		rCategory.GET("", category.ListCategory)
		rCategory.DELETE("/:id", category.DeleteCategory)
		rCategory.GET("/:id", category.GetCategoryById)
	}

	r.Run(utils.HttpPort)
}
