package services

import (
	"app/models"
	"app/ogen"
	"app/pkg/db"
	"context"
	"errors"
)

type MicropostService struct{}

func (s *MicropostService) APIMicropostsCountGet(ctx context.Context) (ogen.APIMicropostsCountGetRes, error) {
	if db.DB == nil {
		return nil, errors.New("データベース接続が利用できません")
	}
	var count int64
	if err := db.DB.WithContext(ctx).Model(&models.Micropost{}).Count(&count).Error; err != nil {
		return &ogen.APIMicropostsCountGetInternalServerError{}, err
	}
	return &ogen.CountResponse{Count: ogen.NewOptInt(int(count))}, nil
}
