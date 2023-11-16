package services

import (
	"app/models"
	"app/ogen"
	"app/pkg/db"
	"context"
	"errors"
	"fmt"
)

type UserService struct{}

func (s *UserService) APIUsersCountGet(ctx context.Context) (ogen.APIUsersCountGetRes, error) {
	if db.DB == nil {
		return nil, errors.New("データベース接続が利用できません")
	}
	var count int64
	if err := db.DB.WithContext(ctx).Model(&models.User{}).Count(&count).Error; err != nil {
		return &ogen.APIUsersCountGetInternalServerError{}, fmt.Errorf("データベースからのユーザー数の取得に失敗しました: %w", err)
	}
	return &ogen.CountResponse{Count: ogen.NewOptInt(int(count))}, nil
}
