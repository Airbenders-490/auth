package database

import (
	"fmt"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"mocklogin/model"
	"os"
)

// DatabaseConnection is a global var initialized as a pointer of type gorm.DB, to be able to pass connection to another package
var DatabaseConnection *gorm.DB

// Connect function connects to the mock database
func Connect() {

	dsn := fmt.Sprintf("user=%s password=%s dbname=%s port=%s sslmode=disable",
		os.Getenv("DB_USERNAME"), os.Getenv("DB_PASSWORD"), os.Getenv("DB_NAME"), os.Getenv("DB_PORT"))

	connection, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		panic("Could not connect to the database")
	}

	DatabaseConnection = connection
	fmt.Println("Connected to the database!")

	// creates database user table given the User type
	if err := connection.AutoMigrate(&model.User{}) ; err != nil {
		panic("Could not auto migrate user table")
	}
}