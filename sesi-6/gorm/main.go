package main

import (
	"fmt"
	"gorm/database"
	"gorm/models"
)

func main() {
	database.StartDB()

	createUser("fitri@mail.com")
}

func createUser(email string) {
	db := database.GetDB()
	if db == nil {
		fmt.Println("Error: Database connection is nil")
		return
	}

	user := models.User{
		Email: email,
	}

	err := db.Create(&user).Error
	if err != nil {
		fmt.Println("Error creating user data: ", err)
		return
	}

	fmt.Println("New user data", user)
}

func getUserById(id uint) {
	if db == nil {
		fmt.Println("Error: Database connection is nil")
		return
	}
}
