package models

type Spot struct {
	ID         int     // 自動採番される一意のID
	Name       string  // スポットの名前(例：伏見稲荷大社)
	VideoTitle string  // 動画のタイトル
	VideoID    string  // Youtube動画URLのID(dQw4w9WgXcQみたいな)
	Prefecture string  // 都道府県名
	Latitude   float64 // 緯度(地図用）
	Longitude  float64 // 経度(地図用）
	TimeSec    int     // 動画中の開始時間
	Note       string  // メモ
}
