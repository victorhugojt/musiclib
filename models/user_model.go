package models

type User struct {
	Id       string
	UserName string
	FullName string
	Email    string
}

type UserService interface {
	CreateUser() error
	GetAllUser() ([]*User, error)
}
