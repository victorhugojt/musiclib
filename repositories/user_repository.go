package repositories

import (
	"fmt"
	"log"

	"musiclib.com.co/musiclib/db"
	"musiclib.com.co/musiclib/models"
)

func AddUser(user *models.User) error {
	InsertUser(user.Id, user.UserName, user.FullName, user.Created_By)
	return nil
}

func GetAllUsers() ([]*models.User, error) {
	users, err := queryUsers()
	if err != nil {
		return nil, fmt.Errorf("error querying users: %w", err)
	}
	return users, nil
}

func queryUsers() ([]*models.User, error) {
	var users []*models.User

	queryStr := `SELECT full_name, email, created_by, created_on FROM users`
	rows, err := db.DB.Query(queryStr)
	if err != nil {
		return nil, fmt.Errorf("error executing query: %w", err)
	}
	defer rows.Close()

	for rows.Next() {
		var user models.User
		err := rows.Scan(&user.FullName, &user.Email, &user.Created_By, &user.Created_On)
		if err != nil {
			return nil, fmt.Errorf("error scanning user: %w", err)
		}
		users = append(users, &user)
	}

	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("error iterating over rows: %w", err)
	}

	return users, nil
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
