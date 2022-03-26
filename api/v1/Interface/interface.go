// Provide the interface for the API
package apiinterface

import (
	"github.com/gin-gonic/gin"
)

type UserInterface interface {
	AddUser(c *gin.Context)
	GetUserById(c *gin.Context)
	ListUsers(c *gin.Context)
	DeleteUser(c *gin.Context)
	UpdateUserById(c *gin.Context)
}

type ArticleInterface interface {
	GetArticleById(c *gin.Context)
	GetArticlesList(c *gin.Context)
	DeleteArticle(c *gin.Context)
	AddArticle(c *gin.Context)
	UpdateArticleById(c *gin.Context)
}

type CategoryInterface interface {
	AddCategory(c *gin.Context)
	ListCategories(c *gin.Context)
	DeleteCategory(c *gin.Context)
	GetCategoryById(c *gin.Context)
}

type TagInterface interface {
	AddTag(c *gin.Context)
	ListTags(c *gin.Context)
	DeleteTag(c *gin.Context)
	GetTagById(c *gin.Context)
}

type MenuInterface interface {
	AddMenu(c *gin.Context)
	ListMenus(c *gin.Context)
	DeleteMenu(c *gin.Context)
	GetMenuById(c *gin.Context)
	UpdateMenuById(c *gin.Context)
}

type RoleInterface interface {
	AddRole(c *gin.Context)
	ListRoles(c *gin.Context)
	DeleteRole(c *gin.Context)
	GetRoleById(c *gin.Context)
	UpdateRole(c *gin.Context)
}

type SdInterface interface {
	HealthCheck(c *gin.Context)
	DiskCheck(c *gin.Context)
	CPUCheck(c *gin.Context)
	RAMCheck(c *gin.Context)
	NetCheck(c *gin.Context)
	HostCheck(c *gin.Context)
}


