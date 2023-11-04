package handlers

import (
	"app/models"
	"app/pkg/db"
	"app/pkg/utils"
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

	utils.WriteJSONResponse(w, http.StatusOK, response)
}
