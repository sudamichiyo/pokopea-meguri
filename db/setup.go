package db

import (
	"database/sql"
	"log"

	_ "github.com/mattn/go-sqlite3" // SQLiteドライバ
)

func InitDB(filepath string) *sql.DB {
	db, err := sql.Open("sqlite3", filepath)
	if err != nil {
		log.Fatal(err)
	}

	if db == nil {
		log.Fatal("DB is nil")
	}

	return db
}

func CreateTable(db *sql.DB) {
	createTableSQL := `CREATE TABLE IF NOT EXISTS spots (
        id INTEGER PRIMARY KEY AUTOINCREMENT,
        name TEXT,
        video_title TEXT,
        video_id TEXT,
        prefecture TEXT,
        latitude REAL,
        longitude REAL,
        time_sec INTEGER,
        note TEXT,
		UNIQUE(name, video_id)
    );`

	stm, err := db.Prepare(createTableSQL)
	if err != nil {
		log.Fatal(err)
	}

	_, err = stm.Exec()
	if err != nil {
		log.Fatal(err)
	}

	log.Println("table creation complete or already exists")
}
