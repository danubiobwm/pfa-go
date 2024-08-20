package database

import (
	"database/sql"
	"log"
)

func InitializeDB(db *sql.DB) error {
	query := `
    CREATE TABLE IF NOT EXISTS orders (
        id VARCHAR(255) PRIMARY KEY,
        price DECIMAL(10,2) NOT NULL,
        tax DECIMAL(10,2) NOT NULL,
        final_price DECIMAL(10,2) NOT NULL
    );`

	_, err := db.Exec(query)
	if err != nil {
		log.Fatalf("Could not create table: %v", err)
		return err
	}

	log.Println("Database table created successfully or already exists.")
	return nil
}
