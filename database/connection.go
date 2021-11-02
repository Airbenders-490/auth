package database

import (
	"fmt"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"mocklogin/model"
	"os"
)

// global var initialized as a pointer of type gorm.DB, to be able to pass connection to another package
var DatabaseConnection *gorm.DB

func Connect() {

	dsn := fmt.Sprintf("user=%s password=%s dbname=%s port=%s sslmode=disable",
		os.Getenv("DB_USERNAME"), os.Getenv("DB_PASSWORD"), os.Getenv("DB_NAME"), os.Getenv("DB_PORT"))

	connection, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		panic("Could not connect to the database")
	}

	DatabaseConnection = connection
	fmt.Println("Connected to the database!")

	connection.AutoMigrate(&model.User{}) // creates database user table given the User type
}