package models

type User struct {
	Id       string `json:"id"`
	UserName string `json:"user_name"`
	FullName string `json:"full_name"`
	Email    string `json:"email"`
}

func NewUser(id, userName, fullName, email string) *User {
	return &User{
		Id:       id,
		UserName: userName,
		FullName: fullName,
		Email:    email,
	}
}
