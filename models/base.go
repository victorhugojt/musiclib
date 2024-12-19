package models

import "time"

type Base struct {
	Id         string    `json:"id"`
	Created_By string    `json:"created_by"`
	Created_On time.Time `json:"created_on"`
}
