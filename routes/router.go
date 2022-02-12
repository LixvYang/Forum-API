package routes

import (
	"context"
	"mixindev/api/v1/article"
	"mixindev/api/v1/category"
	"mixindev/api/v1/menu"
	"mixindev/api/v1/role"
	"mixindev/api/v1/tag"
	"mixindev/api/v1/user"
	"mixindev/middleware"
	"mixindev/model"
	"mixindev/api/sd"
	"net/http"

	_ "mixindev/docs" // docs is generated by Swag CLI, you have to import it.

	"github.com/swaggo/gin-swagger"
	"github.com/swaggo/gin-swagger/swaggerFiles"
	"github.com/gin-gonic/gin"
)
// var tagHandler *tag.TagHandler
// var userHandler *user.UserHandler


func InitRouter(r *gin.Engine, mw ...gin.HandlerFunc) *gin.Engine {

	ctx := context.Background()
	// Middlewares
	r.Use(gin.Recovery())
	r.Use(middleware.NoCache)
	r.Use(middleware.Options)
	r.Use(middleware.Secure)
	r.Use(mw...)
	// 404 Handler
	go r.NoRoute(func(c *gin.Context) {
		c.String(http.StatusNotFound, "The incorrect API router.")
	})

	//swagger api Document
	go r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	go r.POST("v1/login", user.Login)

	go r.GET("v1/auth/userinfo", user.GetUseInfo)
	// r.Use(middleware.AuthMiddleware())

	userHandler := user.NewUserHandler(ctx, model.RedisClient)
	rUser := r.Group("v1/user")
	//auth 
	// rUser.Use(middleware.AuthMiddleware())
	{
		// 登录鉴权后获取用户信息
		go rUser.POST("/add", userHandler.AddUser)
		go rUser.GET("/:id", userHandler.GetUserById)
		go rUser.GET("", userHandler.ListUsers)
		go rUser.DELETE("/:id", userHandler.DeleteUser)
		go rUser.PUT("/:id", userHandler.UpdateUserById)
	}
	
	//文章相关
	articleHandler := article.NewArticleHandler(ctx, model.RedisClient)
	rArticle := r.Group("v1/article")
	// rArticle.Use(middleware.AuthMiddleware())
	{
		go rArticle.GET("/:id", articleHandler.GetArticleById)
		go rArticle.GET("", articleHandler.GetArticlesList)
		go rArticle.DELETE("/:id", articleHandler.DeleteArticle)
		go rArticle.POST("/add", articleHandler.AddArticle)
		go rArticle.PUT(":id",articleHandler.UpdateArticleById)
	}
	
	categoryHandler := category.NewCategoryHandler(ctx, model.RedisClient)
	rCategory := r.Group("v1/category")
	// rCategory.Use(middleware.AuthMiddleware()
	{
		go rCategory.POST("/add", categoryHandler.AddCategory)
		go rCategory.GET("", categoryHandler.ListCategories)
		go rCategory.DELETE("/:id", categoryHandler.DeleteCategory)
		go rCategory.GET("/:id", categoryHandler.GetCategoryById)
	}

	tagHandler := tag.NewTagHandler(ctx, model.RedisClient)
	rTag := r.Group("v1/tag")
	// rTag.Use(middleware.AuthMiddleware()
	{
		rTag.POST("/add", tagHandler.AddTag)
		rTag.GET("", tagHandler.ListTags)
		rTag.DELETE("/:id", tagHandler.DeleteTag)
		rTag.GET("/:id", tagHandler.GetTagById)
	}

	menuHandler := menu.NewMenuHandler(ctx, model.RedisClient)
	rMenu := r.Group("v1/menu")
	// rMenu.Use(middleware.AuthMiddleware()
	{
		rMenu.POST("/add", menuHandler.AddMenu)
		rMenu.GET("", menuHandler.ListMenus)
		rMenu.DELETE("/:id", menuHandler.DeleteMenu)
		rMenu.GET("/:id", menuHandler.GetMenuById)
		rMenu.PUT("/:id", menuHandler.UpdateMenuById)
	}

	roleHandler := role.NewRoleHandler(ctx, model.RedisClient)
	rRole := r.Group("v1/role")
	// rRole.Use(middleware.AuthMiddleware()
	{
		rRole.POST("/add", roleHandler.AddRole)
		rRole.GET("", roleHandler.ListRoles)
		rRole.DELETE("/:id", roleHandler.DeleteRole)
		rRole.GET("/:id", roleHandler.GetRoleById)
		rRole.PUT("/:id", roleHandler.UpdateRole)
	}

	rSd := r.Group("sd")
	{
		rSd.GET("/health", sd.HealthCheck)
		rSd.GET("/disk", sd.DiskCheck)
		rSd.GET("/cpu", sd.CPUCheck)
		rSd.GET("/ram", sd.RAMCheck)
		rSd.GET("net",sd.NetCheck)
		rSd.GET("/host",sd.HostCheck)
	}
	//Will run https.
	// r.RunTLS(":433","./certs/localhost.crt", "./certs/localhost.key")
	return r
}
