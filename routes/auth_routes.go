package routes

import (
	"main/controller"

	"github.com/gin-gonic/gin"
)

// AuthRoutes defines authentication-related routes
func AuthRoutes(router *gin.Engine) {
    auth := router.Group("/api/auth")
    {
        auth.POST("/login", controller.LoginHandler)
    }
}
