package services

import (
	"musiclib.com.co/musiclib/models"
	"musiclib.com.co/musiclib/repositories"
)

func CreateUser(u *models.User) error {
	return repositories.AddUser(u)
}

func GetAllUser() ([]*models.User, error) {
	return repositories.GetAllUsers()
}
