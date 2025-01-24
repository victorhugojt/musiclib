package repositories

import (
	"fmt"

	"musiclib.com.co/musiclib/db"
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

func GetLibraryById(id string) (*models.Library, error) {

	queryStr := `SELECT name, description, user_id FROM libraries`
	rows, err := db.DB.Query(queryStr)
	if err != nil {
		return nil, fmt.Errorf("error executing query: %w", err)
	}
	defer rows.Close()

	var library models.Library
	err = rows.Scan(&library.Name, &library.Description, &library.User_Id)
	if err != nil {
		return nil, fmt.Errorf("error scanning library: %w", err)
	}

	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("error gettingo row: %w", err)
	}

	return &library, nil
}
