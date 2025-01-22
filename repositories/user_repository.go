package repositories

import (
	"fmt"
	"log"

	"musiclib.com.co/musiclib/db"
	"musiclib.com.co/musiclib/models"
)

var userList []*models.User

func AddUser(user *models.User) error {
	InsertUser(user.Id, user.UserName, user.FullName, user.Created_By)
	userList = append(userList, user)
	return nil
}

func GetAllUsers() ([]*models.User, error) {
	return userList, nil
}

func InsertUser(fullName, email, createdBy, createdOn string) {
	insertUserSQL := `INSERT INTO users (full_name, email, created_by, created_on) VALUES (?, ?, ?, ?)`
	statement, err := db.DB.Prepare(insertUserSQL)
	if err != nil {
		log.Fatalf("Error preparing statement: %v", err)
	}
	defer statement.Close()

	_, err = statement.Exec(fullName, email, createdBy, createdOn)
	if err != nil {
		log.Fatalf("Error executing statement: %v", err)
	}

	fmt.Println("User inserted successfully")
}
