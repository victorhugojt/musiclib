package repositories

import (
	"fmt"
	"log"
	"time"

	"musiclib.com.co/musiclib/db"
	"musiclib.com.co/musiclib/models"
	"musiclib.com.co/musiclib/utils"
)

func GetUserByEmail(email string) (*models.User, error) {

	queryStr := `SELECT id, user_name, full_name, email, password FROM users WHERE email = ?`
	row := db.DB.QueryRow(queryStr, email)

	var user models.User
	err := row.Scan(&user.Id, &user.UserName, &user.FullName, &user.Email, &user.Password)
	if err != nil {
		return nil, fmt.Errorf("error scanning User: %w", err)
	}

	return &user, nil
}

func AddUser(user *models.User) error {
	InsertUser(user.Id, user.FullName, user.Email, user.Created_By, user.Created_On, user.UserName, user.Password)
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

	queryStr := `SELECT id, user_name, full_name, email, created_by, created_on FROM users`
	rows, err := db.DB.Query(queryStr)
	if err != nil {
		return nil, fmt.Errorf("error executing query: %w", err)
	}
	defer rows.Close()

	for rows.Next() {
		var user models.User
		err := rows.Scan(&user.Id, &user.UserName, &user.FullName, &user.Email, &user.Created_By, &user.Created_On)
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

func InsertUser(id, fullName, email, createdBy string, createdOn time.Time, user_name, password string) {

	hashedPassword, err := utils.HashPassword(password)

	insertUserSQL := `INSERT INTO users (id, full_name, email, created_by, created_on, user_name, password) VALUES (?, ?, ?, ?, ?, ?, ?)`
	statement, err := db.DB.Prepare(insertUserSQL)
	if err != nil {
		log.Fatalf("Error preparing statement: %v", err)
	}
	defer statement.Close()

	_, err = statement.Exec(id, fullName, email, createdBy, createdOn, user_name, hashedPassword)
	if err != nil {
		log.Fatalf("Error executing statement: %v", err)
	}

	fmt.Println("User inserted successfully")
}
