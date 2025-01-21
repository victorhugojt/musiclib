package models

import (
	"time"

	"github.com/google/uuid"
)

type BaseTemplate struct {
	Id         string    `json:"id"`
	Created_On time.Time `json:"created_on"`
}

func GenerateUUID() string {
	return uuid.New().String()
}

func NewBaseTemplate(base BaseTemplate) *BaseTemplate {
	return &BaseTemplate{
		Id:         GenerateUUID(),
		Created_On: time.Now(),
	}
}
