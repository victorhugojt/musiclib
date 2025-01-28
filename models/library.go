package models

import "fmt"

type Library struct {
	Name        string `json:"name"`
	Description string `json:"description"`
	Created_By  string `json:"created_by"`
	User_Id     string `json:"user_id"`
	BaseTemplate
}

func (l Library) String() string {
	return fmt.Sprintf("%v - %v - %v - %v - %v", l.Id, l.Name, l.Description, l.Created_By, l.Created_On)
}
