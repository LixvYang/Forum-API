package routes

import (
	"mixindev/api/v1/article"
	"mixindev/api/v1/category"
	"mixindev/api/v1/tag"
	"mixindev/api/v1/menu"
	"mixindev/api/v1/user"
	"mixindev/api/v1/role"
	"mixindev/utils"

	"github.com/gin-gonic/gin"
)

func InitRouter() {
	gin.SetMode(utils.AppMode)
	r := gin.Default()

	r.POST("v1/login", user.Login)

	r.GET("v1/auth/userinfo",user.GetUseInfo)
	rUser := r.Group("v1/user")
	{
		rUser.POST("", user.AddUser)
		rUser.GET("/:id", user.GetUserById)
		rUser.GET("", user.ListUsers)
		rUser.DELETE("/:id", user.DeleteUser)
		rUser.PUT("/:id", user.UpdateUserById)
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
		rCategory.GET("", category.ListCategories)
		rCategory.DELETE("/:id", category.DeleteCategory)
		rCategory.GET("/:id", category.GetCategoryById)
	}

	rTag := r.Group("v1/tag")
	{
		rTag.POST("", tag.AddTag)
		rTag.GET("", tag.ListTags)
		rTag.DELETE("/:id", tag.DeleteTag)
		rTag.GET("/:id", tag.GetTagById)
	}

	rMenu := r.Group("v1/menu")
	{
		rMenu.POST("", menu.AddMenu)
		rMenu.GET("", menu.ListMenus)
		rMenu.DELETE("/:id", menu.DeleteMenu)
		rMenu.GET("/:id", menu.GetMenuById)
		rMenu.PUT("/:id", menu.UpdateMenuById)
	}

	rRole := r.Group("v1/role")
	{
		rRole.POST("", role.AddRole)
		rRole.GET("", role.ListRoles)
		rRole.DELETE("/:id", role.DeleteRole)
		rRole.GET("/:id", role.GetRoleById)
		rRole.PUT("/:id", role.UpdateRole)
	}

	r.Run(utils.HttpPort)
}
