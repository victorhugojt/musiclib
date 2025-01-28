package repositories

import (
	"fmt"
	"log"
	"time"

	"musiclib.com.co/musiclib/db"
	"musiclib.com.co/musiclib/models"
)

func AddLibrary(library *models.Library) error {
	fmt.Print(library)
	InsertLibrary(library.Id, library.Name, library.Description, library.Created_By, library.User_Id, library.Created_On)
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

	queryStr := `SELECT id, name, description, user_id FROM libraries WHERE id = ?`
	row := db.DB.QueryRow(queryStr, id)

	var library models.Library
	err := row.Scan(&library.Id, &library.Name, &library.Description, &library.User_Id)
	if err != nil {
		return nil, fmt.Errorf("error scanning library: %w", err)
	}

	return &library, nil
}

func InsertLibrary(id, name, description, createdBy, user_id string, created_on time.Time) {
	insertUserSQL := `INSERT INTO libraries (id, name, description, created_by, user_id, created_on) VALUES (?, ?, ?, ?, ?, ?)`
	statement, err := db.DB.Prepare(insertUserSQL)
	if err != nil {
		log.Fatalf("Error preparing statement: %v", err)
	}
	defer statement.Close()

	_, err = statement.Exec(id, name, description, createdBy, user_id, created_on)
	if err != nil {
		log.Fatalf("Error inserting library: %v", err)
	}

	fmt.Println("Library inserted successfully")
}

func queryLibraries() ([]*models.Library, error) {
	var libraries []*models.Library

	queryStr := `SELECT id, name, description, user_id FROM libraries`
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
