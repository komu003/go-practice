package handlers

import (
	"app/models"
	"app/pkg/db"
	"encoding/json"
	"net/http"
)

type MicropostsCountResponse struct {
	Count int `json:"count"`
}

func MicropostsCountHandler(w http.ResponseWriter, r *http.Request) {
	var count int64

	if err := db.DB.Model(&models.Micropost{}).Count(&count).Error; err != nil {
		http.Error(w, "データベースのエラー", http.StatusInternalServerError)
		return
	}

	response := MicropostsCountResponse{Count: int(count)}

	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "http://localhost:3000")
	json.NewEncoder(w).Encode(response)
}
