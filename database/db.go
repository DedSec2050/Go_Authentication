package database

import (
	"log"
	"main/models"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var DB *gorm.DB

// ConnectDatabase initializes the database connection
func ConnectDatabase() {
    db, err := gorm.Open(sqlite.Open("localdb.sqlite"), &gorm.Config{})
    if err != nil {
        log.Fatalf("Failed to connect to database: %v", err)
    }

    // Assign to the global variable
    DB = db

    // Migrate the User model
    if err := db.AutoMigrate(&models.User{}); err != nil {
        log.Fatalf("Failed to migrate database: %v", err)
    }
    log.Println("Database connected and migrated")
	logAllContents(db)
}

func logAllContents(db *gorm.DB) {
	var users []models.User

	if err := db.Find(&users).Error; err != nil {
		log.Printf("Failed to find users: %v", err)
		return
	}
	db.Find(&users)
	for _, user := range users {
		log.Printf("User: %v\n", user)
	}
}