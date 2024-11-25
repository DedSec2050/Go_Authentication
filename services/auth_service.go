package services

import (
	"errors"
	"fmt"
	"main/repositories"
	"main/utils"
)

// LoginUser authenticates the user and generates a JWT if successful
func LoginUser(email, password string) (string, error) {
    // Find user by email
    user, err := repositories.FindUserByEmail(email)
    if err != nil {
		fmt.Print("User not Found")
        return "", errors.New("user not found")
    }

    // Verify the password
    if err := utils.VerifyPassword(user.Password, password); err != nil {
		fmt.Println("Password Verification Failed")
        return "", errors.New("invalid password")
    }

    // Generate JWT token
	fmt.Println("Password Verified !! Generating JWT")
    return utils.GenerateJWT(user.ID)
}
