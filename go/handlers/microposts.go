package handlers

import (
	"app/models"
	"app/pkg/db"
	"app/pkg/utils"
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

	utils.WriteJSONResponse(w, http.StatusOK, response)
}
