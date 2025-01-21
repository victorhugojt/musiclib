package models

import "fmt"

type User struct {
	UserName   string `json:"user_name"`
	FullName   string `json:"full_name"`
	Email      string `json:"email"`
	Created_By string `json:"created_by"`
	BaseTemplate
}

func NewUser(userName, fullName, email string, created_by string) *User {
	defaultBaseTemplate := BaseTemplate{}
	baseTemplate := NewBaseTemplate(defaultBaseTemplate)

	return &User{
		UserName:     userName,
		FullName:     fullName,
		Email:        email,
		Created_By:   created_by,
		BaseTemplate: *baseTemplate,
	}
}

func (u User) String() string {
	return fmt.Sprintf("%v %v", u.UserName, u.FullName)
}
