package repositories

import (
	"musiclib.com.co/musiclib/models"
)

// userList stores all users in memory
var userList []*models.User

// AddUser adds a new user to the in-memory storage
func AddUser(user *models.User) error {
	userList = append(userList, user)
	return nil
}

// GetAllUsers returns all users from the in-memory storage
func GetAllUsers() ([]*models.User, error) {
	return userList, nil
}
