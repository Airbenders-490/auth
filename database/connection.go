package database

import (
	"fmt"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"mocklogin/model"
)

// global var initialized as a pointer of type gorm.DB, to be able to pass connection to another package
var DatabaseConnection *gorm.DB

func Connect() {
	dsn := "user=postgres password=password dbname=postgres port=62626 sslmode=disable"

	connection, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		panic("Could not connect to the database")
	}

	DatabaseConnection = connection
	fmt.Println("Connected to the database!")

	connection.AutoMigrate(&model.User{}) // creates database user table given the User type
}