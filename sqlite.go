package gamewallpaper

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/mattn/go-sqlite3"
)

var db *sql.DB

const dbPath = "yostar-gallery.db"

func initDB() error {
	var err error
	db, err = sql.Open("sqlite3", dbPath)
	if err != nil {
		return fmt.Errorf("failed to open database: %w", err)
	}

	createTable := `
		CREATE TABLE IF NOT EXISTS yostar_gallery (
			id INTEGER PRIMARY KEY AUTOINCREMENT,
			id_gallery VARCHAR(255) NOT NULL,
			game VARCHAR(255) NOT NULL,
			type VARCHAR(255) NOT NULL,
			file_name VARCHAR(255) NOT NULL,
			url VARCHAR(255) NOT NULL,
			created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
		);
	`
	_, err = db.Exec(createTable)
	if err != nil {
		db.Close()
		return fmt.Errorf("failed to create table: %w", err)
	}
	fmt.Println("=======DB created=======")
	return nil
}

func GetSqliteDb() *sql.DB {
	if db == nil {
		if err := initDB(); err != nil {
			log.Fatalf("Database initialization failed: %v", err)
		}
	}
	return db
}
