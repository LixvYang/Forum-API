package main

import (
	"mixindev/model"
	"mixindev/routes"
)

func main() {
	model.InitDb()
	routes.InitRouter()
}
