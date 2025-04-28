package models

type Spot struct {
	ID         int     `json:"id"`          // 自動採番される一意のID
	Name       string  `json:"name"`        // スポットの名前(例：伏見稲荷大社)
	VideoTitle string  `json:"video_title"` // 動画のタイトル
	VideoID    string  `json:"video_id"`    // Youtube動画URLのID(dQw4w9WgXcQみたいな)
	Prefecture string  `json:"prefecture"`  // 都道府県名
	Latitude   float64 `json:"latitude"`    // 緯度(地図用）
	Longitude  float64 `json:"longitude"`   // 経度(地図用）
	TimeSec    int     `json:"time_sec"`    // 動画中の開始時間
	Note       string  `json:"note"`        // メモ
}
