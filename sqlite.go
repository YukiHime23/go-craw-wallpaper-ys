package crawal

import (
	"database/sql"
	"log"

	_ "github.com/mattn/go-sqlite3"
)

var db *sql.DB

func init() {
	var err error
	// Kết nối đến cơ sở dữ liệu SQLite
	db, err = sql.Open("sqlite3", "data-azur-lane.db")
	if err != nil {
		log.Fatal(err)
	}
}

func GetSqliteDb() *sql.DB {
	return db
}
