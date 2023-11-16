package services

import (
	"app/models"
	"app/ogen"
	"app/pkg/db"
	"context"
	"fmt"
)

type UserService struct{}

func (s *UserService) APIUsersCountGet(ctx context.Context) (ogen.APIUsersCountGetRes, error) {
	var count int64
	if err := db.DB.WithContext(ctx).Model(&models.User{}).Count(&count).Error; err != nil {
		errMsg := fmt.Errorf("データベースからのユーザー数の取得に失敗しました: %w", err)
		return &ogen.APIUsersCountGetInternalServerError{}, errMsg
	}
	return &ogen.CountResponse{Count: ogen.NewOptInt(int(count))}, nil
}
