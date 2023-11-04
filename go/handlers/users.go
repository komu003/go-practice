package handlers

import (
	"app/models"
	"app/pkg/db"
	"encoding/json"
	"net/http"
)

type UsersCountResponse struct {
	Count int `json:"count"`
}

func UsersCountHandler(w http.ResponseWriter, r *http.Request) {
	var count int64

	if err := db.DB.Model(&models.User{}).Count(&count).Error; err != nil {
		http.Error(w, "データベースのエラー", http.StatusInternalServerError)
		return
	}

	response := UsersCountResponse{Count: int(count)}

	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "http://localhost:3000")
	json.NewEncoder(w).Encode(response)
}
