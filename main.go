package main

import (
	"mixindev/middleware"
	"mixindev/model"
	"mixindev/routes"
	"mixindev/utils"

	"github.com/gin-gonic/gin"
)

func main() {

	model.InitDb()
	// routes.InitRouter()

	// Set gin mode
	gin.SetMode(utils.AppMode)

	//Create gin engine
	g := gin.New()

	//Routes
	routes.InitRouter(
		// Cores.
		g,
		// Middlwares.
		middleware.Log(),
		middleware.RequestId(),
	)

	g.Run(utils.HttpPort)
}
