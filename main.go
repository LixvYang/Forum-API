package main

import (
	"mixindev/middleware"
	"mixindev/model"
	"mixindev/routes"
	"mixindev/utils"

	"github.com/gin-gonic/gin"
)


// @title Forum-API
// @version 1.0
// @description mixindev API doc
// @securityDefinitions.apikey ApiKeyAuth
// @in header
// @name Authorization
// @contact.name Lixv
// @contact.url http://www.swagger.io/support
// @contact.email yanglixin621@gmail.com
// @host   localhost:3000
// @BasePath
func main() {

	go model.InitDb()
	// routes.InitRouter()

	// Set gin mode
	go gin.SetMode(utils.AppMode)

	//Create gin engine
	g := gin.New()

	//Routes
	go routes.InitRouter(
		// Cores.
		g,
		// Middlwares.
		middleware.Log(),
		middleware.RequestId(),
	)

	g.Run(utils.HttpPort)
}
