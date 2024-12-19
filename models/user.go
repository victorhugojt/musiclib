package models

import "time"

type User struct {
	UserName string `json:"user_name"`
	FullName string `json:"full_name"`
	Email    string `json:"email"`
	Base
}

func NewUser(id, userName, fullName, email string, created_by string) *User {
	return &User{
		UserName: userName,
		FullName: fullName,
		Email:    email,
		Base: Base{
			Id:         id,
			Created_By: created_by,
			Created_On: time.Now(),
		},
	}
}
