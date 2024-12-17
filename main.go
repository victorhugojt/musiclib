package main

import (
	"fmt"

	"musiclib.com.co/musiclib/models"
	"musiclib.com.co/musiclib/services"
	"musiclib.com.co/musiclib/utils"
)

func main() {
	fmt.Println("Hello, Music Library!")
	/*user1 := models.User{
		Id:       "1",
		UserName: "johndoe",
		FullName: "John Doe",
		Email:    "john@example.com",
	}
	services.CreateUser(&user1)*/

	user := models.NewUser("3", "sofi", "sofi Doe", "sofi@example.com")
	services.CreateUser(user)
	utils.Encode(user, "test_file.json")

	/*user2 := models.User{
		Id:       "2",
		UserName: "anita",
		FullName: "Anita Doe",
		Email:    "anita@example.com",
	}
	services.CreateUser(&user2)*/

	users, err := services.GetAllUser()
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	fmt.Println("\nUsers in the system:")
	fmt.Println("--------------------")
	for _, user := range users {
		fmt.Printf("Id: %s\n", user.Id)
		fmt.Printf("Username: %s\n", user.UserName)
		fmt.Printf("Full Name: %s\n", user.FullName)
		fmt.Printf("Email: %s\n", user.Email)
		fmt.Println("--------------------")
	}
}
