package repositories

import (
	"errors"
	"main/database"
	"main/models"
)

// FindUserByEmail retrieves a user by their email
func FindUserByEmail(email string) (*models.User, error) {
    var user models.User
    result := database.DB.Where("email = ?", email).First(&user)
    if result.Error != nil {
        return nil, errors.New("user not found")
    }
    return &user, nil
}
