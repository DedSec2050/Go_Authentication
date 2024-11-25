package controller

import (
	"fmt"
	"main/services"
	"net/http"

	"github.com/gin-gonic/gin"
)

// LoginRequest represents the expected JSON body for login requests
type LoginRequest struct {
    Email    string `json:"email" binding:"required,email"`
    Password string `json:"password" binding:"required,min=8"`
}

// LoginHandler handles user login requests
func LoginHandler(c *gin.Context) {
    var loginReq LoginRequest

    // Parse and validate JSON request body
    if err := c.ShouldBindJSON(&loginReq); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }

    // Authenticate user
    token, err := services.LoginUser(loginReq.Email, loginReq.Password)
    if err != nil {
		fmt.Println("Failed To authenticate User")
        c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid email or password"})
        return
    }

    // Respond with JWT token
    c.JSON(http.StatusOK, gin.H{
        "message": "Login successful",
        "token":   token,
    })
}
