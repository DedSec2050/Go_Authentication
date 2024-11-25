package main

import (
	"main/database"
	"main/routes"

	"github.com/gin-gonic/gin"
)

func main() {
    // Connect to the database
    database.ConnectDatabase()

    // Initialize the Gin router
    router := gin.Default()

    // Register routes
    routes.AuthRoutes(router)

    // Start the server
    router.Run(":8080")
}
