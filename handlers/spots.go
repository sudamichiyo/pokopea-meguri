package handlers

import (
	"database/sql"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/sudamichiyo/pokopea-meguri/models"
)

func getSpotsByPrefecture(db *sql.DB, prefecture string) ([]models.Spot, error) {
	rows, err := db.Query("SELECT * FROM spots WHERE prefecture = ?", prefecture)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var spots []models.Spot
	for rows.Next() {
		var spot models.Spot
		if err := rows.Scan(&spot.ID, &spot.Name, &spot.VideoTitle, &spot.VideoID, &spot.Prefecture, &spot.Latitude, &spot.Longitude, &spot.TimeSec, &spot.Note); err != nil {
			return nil, err
		}
		spots = append(spots, spot)
	}

	return spots, nil
}

func SpotsHandler(db *sql.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		prefecture := c.DefaultQuery("prefecture", "")
		if prefecture == "" {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": "prefecture is required",
			})
			return
		}

		spots, err := getSpotsByPrefecture(db, prefecture)
		if err != nil {
			log.Println("error fetching spots:", err)
			c.JSON(http.StatusInternalServerError, gin.H{
				"error": "internal server error",
			})
			return
		}

		c.JSON(http.StatusOK, spots)
	}
}
