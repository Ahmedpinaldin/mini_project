package Config

import (
	"database/sql"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

func LoadConfig() *sql.DB {
	// koneksi ke database
	db, err := sql.Open("mysql", "root@tcp(localhost:3306)/quiz_app")
	if err != nil {
		log.Fatal("Failed to connect to database:", err)
	}

	// Uji koneksi ke database
	err = db.Ping()
	if err != nil {
		log.Fatal("Failed to ping database:", err)
	}
	return db
}
