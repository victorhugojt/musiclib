package main

import (
	"fmt"

	"musiclib.com.co/musiclib/models"
	"musiclib.com.co/musiclib/rest/routers"
	"musiclib.com.co/musiclib/services"
)

func main() {
	fmt.Println("Hello, Music Library!")

	user0 := models.NewUser("0", "sofi", "sofi Doe", "sofi@example.com", "admin")
	services.CreateUser(user0)

	user1 := models.NewUser("1", "emi", "emi Doe", "emi@example.com", "admin")
	services.CreateUser(user1)

	user2 := models.NewUser("2", "lini", "lini Doe", "lini@example.com", "admin")
	services.CreateUser(user2)

	routers.StartRouting()

}
