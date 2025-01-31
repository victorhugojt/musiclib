package services

import (
	"errors"

	"musiclib.com.co/musiclib/models"
	"musiclib.com.co/musiclib/repositories"
	"musiclib.com.co/musiclib/utils"
)

func CreateUser(u *models.User) (*models.User, error) {
	err := repositories.AddUser(u)
	return u, err
}

func GetAllUser() ([]*models.User, error) {
	return repositories.GetAllUsers()
}

func ValidateCredentials(u *models.User) error {
	user, err := repositories.GetUserByEmail(u.Email)

	if err != nil {
		return err
	}

	passwordValid := utils.CheckPassWord(u.Password, user.Password)

	if !passwordValid {
		return errors.New("credentials are not valid")
	}

	return nil

}
