package repositories

import (
	"fmt"
	"log"

	"musiclib.com.co/musiclib/db"
	"musiclib.com.co/musiclib/models"
)

var libraryList []*models.Library

func AddLibrary(library *models.Library) error {
	InsertLibrary(library.Name, library.Description, library.Created_By, library.User_Id)
	libraryList = append(libraryList, library)
	return nil
}

func GetAllLibraries() ([]*models.Library, error) {
	libraries, err := queryLibraries()
	if err != nil {
		return nil, fmt.Errorf("error querying libraries: %w", err)
	}
	return libraries, nil
}

func GetLibraryById(id string) (*models.Library, error) {

	queryStr := `SELECT name, description, user_id FROM libraries WHERE id = ?`
	row := db.DB.QueryRow(queryStr, id)

	var library models.Library
	err := row.Scan(&library.Id, &library.Name, &library.Description, &library.User_Id)
	if err != nil {
		return nil, fmt.Errorf("error scanning library: %w", err)
	}

	return &library, nil
}

func InsertLibrary(name, description, createdBy, user_id string) {
	insertUserSQL := `INSERT INTO libraries (name, description, created_by, user_id) VALUES (?, ?, ?, ?)`
	statement, err := db.DB.Prepare(insertUserSQL)
	if err != nil {
		log.Fatalf("Error preparing statement: %v", err)
	}
	defer statement.Close()

	_, err = statement.Exec(name, description, createdBy, user_id)
	if err != nil {
		log.Fatalf("Error executing statement: %v", err)
	}

	fmt.Println("Library inserted successfully")
}

func queryLibraries() ([]*models.Library, error) {
	var libraries []*models.Library

	queryStr := `SELECT id, name, description, created_by, user_id FROM libraries`
	rows, err := db.DB.Query(queryStr)
	if err != nil {
		return nil, fmt.Errorf("error executing query: %w", err)
	}
	defer rows.Close()

	for rows.Next() {
		var library models.Library
		err := rows.Scan(&library.Id, &library.Name, &library.Description, &library.User_Id)
		if err != nil {
			return nil, fmt.Errorf("error scanning library: %w", err)
		}
		libraries = append(libraries, &library)
	}

	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("error iterating over rows: %w", err)
	}

	return libraries, nil
}
