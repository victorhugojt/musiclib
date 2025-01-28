package main

import (
	"fmt"

	"musiclib.com.co/musiclib/db"
	"musiclib.com.co/musiclib/rest/routers"
)

func main() {
	fmt.Println("Hello, Music Library!")

	db.InitDB()

	routers.StartRouting()

}
