package main

import (
	"log"

	"github.com/gin-gonic/gin"
	"github.com/sudamichiyo/pokopea-meguri/db"
	"github.com/sudamichiyo/pokopea-meguri/handlers"
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

	// GinのEngineのインスタンスを取得
	r := gin.Default()

	// APIエンドポイントを設定
	r.GET("/spots", handlers.SpotsHandler(database))

	// サーバーを起動
	log.Println("server started at :8080")
	r.Run()
}
