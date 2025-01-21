package models

type Library struct {
	Name        string `json:"name"`
	Description string `json:"description"`
	Created_By  string `json:"created_by"`
	User_Id     string `json:"user_id"`
	BaseTemplate
}
