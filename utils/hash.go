package utils

import (
	"fmt"

	"golang.org/x/crypto/bcrypt"
)

// HashPassword hashes a plaintext password
func HashPassword(password string) (string, error) {
    hash, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	fmt.Println("Hashing Password")
	fmt.Println("Password Hashed : ", string(hash))
	fmt.Println()
    return string(hash), err
}

// VerifyPassword compares a hashed password with a plaintext password
func VerifyPassword(hashedPassword, password string) error {
	
	

	fmt.Println()

	fmt.Println("Verifying Password")
	
	err := bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password))
	if err != nil {
		fmt.Println("Password verification failed \n Reason: ", err)
		return fmt.Errorf("invalid password: %v", err)
	}
	fmt.Println("Password verified successfully")
	return nil
}
