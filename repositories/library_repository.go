package repositories

import (
	"musiclib.com.co/musiclib/models"
)

var libraryList []*models.Library

func AddLibrary(library *models.Library) error {
	libraryList = append(libraryList, library)
	return nil
}

func GetAllLibraries() ([]*models.Library, error) {
	return libraryList, nil
}
