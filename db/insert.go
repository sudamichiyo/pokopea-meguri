package db

import (
	"database/sql"
	"log"
)

// 仮データを挿入する関数
func InsertMockData(db *sql.DB) {
	// OR IGNOREは、指定した条件で既にレコードが存在する場合、その挿入を無視する
	insertSQL := `INSERT OR IGNORE INTO spots (name, video_title, video_id, prefecture, latitude, longitude, time_sec, note)
				  VALUES (?, ?, ?, ?, ?, ?, ?, ?)`

	mockSpots := []struct {
		name, videoTitle, videoID, prefecture, note string
		latitude, longitude                         float64
		timeSec                                     int
	}{
		{
			name:       "伏見稲荷大社",
			videoTitle: "京都のおすすめ観光地TOP5",
			videoID:    "abc123xyz",
			prefecture: "京都",
			latitude:   34.9671,
			longitude:  135.7727,
			timeSec:    85,
			note:       "鳥居のトンネルがすごい！",
		},
		{
			name:       "美ら海水族館",
			videoTitle: "沖縄旅行Vlog",
			videoID:    "xyz789abc",
			prefecture: "沖縄",
			latitude:   26.6944,
			longitude:  127.8778,
			timeSec:    120,
			note:       "ジンベエザメが見られる水族館！",
		},
		{
			name:       "アメリカンビレッジ",
			videoTitle: "【弾丸】ついに南国へ!2泊3日の沖縄旅行へ行ってきました!!【DAY1】",
			videoID:    "F64JNxko1Ws",
			prefecture: "沖縄",
			latitude:   26.6944,
			longitude:  127.8778,
			timeSec:    120,
			note:       "実際に使用するデータはこんな感じ！",
		},
	}

	for _, spot := range mockSpots {
		_, err := db.Exec(insertSQL, spot.name, spot.videoTitle, spot.videoID, spot.prefecture,
			spot.latitude, spot.longitude, spot.timeSec, spot.note)
		if err != nil {
			log.Fatal("temporary data insert error:", err)
		}
	}
	log.Println("temporary data insertion complete!")
}
