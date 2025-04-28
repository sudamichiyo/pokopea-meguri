package main

import (
	"log"

	"github.com/sudamichiyo/pokopea-meguri/db"
)

func main() {
	// DB接続
	database := db.InitDB("./db/spots.db") // SQLiteファイル名
	defer database.Close()

	// spotsテーブル作成
	db.CreateTable(database)

	// 仮データを挿入
	db.InsertMockData(database)

	log.Println("DB setup OK!")
}
