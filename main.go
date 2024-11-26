package main

import (
	"main/database"
	"main/routes"

	"github.com/gin-gonic/gin"
)

func main() {
	database.InitDatabase()

	router := gin.Default()
	routes.AuthRoutes(router)

	router.Run(":8080")
}
