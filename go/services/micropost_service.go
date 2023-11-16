package services

import (
	"app/models"
	"app/ogen"
	"app/pkg/db"
	"context"
	"fmt"
)

type MicropostService struct{}

func (s *MicropostService) APIMicropostsCountGet(ctx context.Context) (ogen.APIMicropostsCountGetRes, error) {
	var count int64
	if err := db.DB.WithContext(ctx).Model(&models.Micropost{}).Count(&count).Error; err != nil {
		errMsg := fmt.Errorf("データベースからのマイクロポスト数の取得に失敗しました: %w", err)
		return &ogen.APIMicropostsCountGetInternalServerError{}, errMsg
	}
	return &ogen.CountResponse{Count: ogen.NewOptInt(int(count))}, nil
}
