package models

type User struct {
	Id       string
	UserName string
	FullName string
	Email    string
}

func NewUser(id, userName, fullName, email string) *User {
	return &User{
		Id:       id,
		UserName: userName,
		FullName: fullName,
		Email:    email,
	}
}