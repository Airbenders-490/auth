package database

import (
	"fmt"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"mocklogin/model"
)

func Connect() {
	dsn := "user=postgres password=password dbname=postgres port=62626 sslmode=disable"

	connection, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		panic("Could not connect to the database")
	}

	fmt.Println("Connected to the database!")

	connection.AutoMigrate(&model.User{})
}