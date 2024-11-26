package routes

import (
	"main/controllers"
	"main/middlewares"

	"github.com/gin-gonic/gin"
)

func AuthRoutes(router *gin.Engine) {
	auth := router.Group("/auth")
	{
		auth.POST("/signup", controllers.SignUp)
		auth.POST("/signin", controllers.SignIn)
	}

	protected := router.Group("/protected")
	protected.Use(middlewares.AuthMiddleware())
	{
		protected.GET("/profile", func(c *gin.Context) {
			userID, _ := c.Get("userID")
			c.JSON(200, gin.H{"userID": userID})
		})
	}
}
