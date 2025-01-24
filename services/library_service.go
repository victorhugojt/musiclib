package services

import (
	"musiclib.com.co/musiclib/models"
	"musiclib.com.co/musiclib/repositories"
)

func CreateLibrary(l *models.Library) (*models.Library, error) {
	err := repositories.AddLibrary(l)
	return l, err
}

func GetAllLibraries() ([]*models.Library, error) {
	return repositories.GetAllLibraries()
}

func GetLibraryById(id string) (*models.Library, error) {
	library, err := repositories.GetLibraryById(id)
	return library, err

}
