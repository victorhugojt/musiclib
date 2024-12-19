package models

import "fmt"

type User struct {
	UserName string `json:"user_name"`
	FullName string `json:"full_name"`
	Email    string `json:"email"`
	BaseTemplate
}

func NewUser(userName, fullName, email string, created_by string) *User {
	defaultBaseTemplate := BaseTemplate{}
	baseTemplate := BaseTemplate.NewBaseTemplate(defaultBaseTemplate, created_by)

	return &User{
		UserName:     userName,
		FullName:     fullName,
		Email:        email,
		BaseTemplate: *baseTemplate,
	}
}

func (u User) String() string {
	return fmt.Sprintf("%v %v", u.UserName, u.FullName)
}
