package models

import (
	"time"

	"github.com/google/uuid"
)

type BaseTemplate struct {
	Id         string    `json:"id"`
	Created_By string    `json:"created_by"`
	Created_On time.Time `json:"created_on"`
}

func GenerateUUID() string {
	return uuid.New().String()
}

func (base BaseTemplate) NewBaseTemplate(created_by string) *BaseTemplate {
	return &BaseTemplate{
		Id:         GenerateUUID(),
		Created_By: created_by,
		Created_On: time.Now(),
	}
}
