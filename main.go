package main

import (
	"log"

	"github.com/sudamichiyo/pokopea-meguri/db"
)

func main() {
	database := db.InitDB("db/spots.db") // SQLiteファイル名
	db.CreateTable(database)
	log.Println("DB setup OK!")
}
