package utils

import (
	"encoding/json"
	"os"

	"musiclib.com.co/musiclib/models"
)

func Encode(user *models.User, fileName string) error {

	json, error := json.Marshal(user)

	if error != nil {
		return error
	}

	os.WriteFile(fileName, json, 0644)

	return nil

}
