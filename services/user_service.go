package services

import (
	"musiclib.com.co/musiclib/models"
	"musiclib.com.co/musiclib/repositories"
)

func CreateUser(u *models.User) (*models.User, error) {
	err := repositories.AddUser(u)
	return u, err
}

func GetAllUser() ([]*models.User, error) {
	return repositories.GetAllUsers()
}
