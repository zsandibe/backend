package db

import (
	"backend/internal/models"
	"database/sql"
	"log"

	_ "github.com/mattn/go-sqlite3"
)

var DB *sql.DB

func InitDB() error {
	db, err := sql.Open("sqlite3", "./data.db")
	if err != nil {
		return err
	}
	DB = db

	err = createTables()
	if err != nil {
		return err
	}

	return nil
}

func createTables() error {
	_, err := DB.Exec(`
		CREATE TABLE IF NOT EXISTS organizations (
			id INTEGER PRIMARY KEY AUTOINCREMENT,
			bin TEXT,
			name TEXT,
			address TEXT,
			workers INTEGER
		);

		CREATE TABLE IF NOT EXISTS projects (
			id INTEGER PRIMARY KEY AUTOINCREMENT,
			name TEXT,
			due_date TEXT,
			owner INTEGER,
			participants TEXT
		);
	`)
	if err != nil {
		log.Println("Error creating tables:", err)
		return err
	}

	return nil
}

func GetOrganizations() ([]models.Organization, error) {
	query := "SELECT id, bin, name, address, workers FROM organizations"
	rows, err := DB.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var organizations []models.Organization
	for rows.Next() {
		var organization models.Organization
		err := rows.Scan(&organization.ID, &organization.BIN, &organization.Name, &organization.Address, &organization.Workers)
		if err != nil {
			return nil, err
		}
		organizations = append(organizations, organization)
	}

	err = rows.Err()
	if err != nil {
		return nil, err
	}

	return organizations, nil
}

func AddOrganization(bin, name, address string, workers int) error {
	query := "INSERT INTO organizations (bin, name, address, workers) VALUES (?, ?, ?, ?)"
	_, err := DB.Exec(query, bin, name, address, workers)
	if err != nil {
		return err
	}

	return nil
}
