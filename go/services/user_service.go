package services

import (
	"app/models"
	"app/ogen"
	"app/pkg/db"
	"context"
	"fmt"
)

type UserService struct{}

func (s *UserService) APIUsersGet(ctx context.Context) (ogen.APIUsersGetRes, error) {
	var users []models.User
	if err := db.DB.WithContext(ctx).Find(&users).Error; err != nil {
		errMsg := fmt.Errorf("データベースからのユーザー一覧の取得に失敗しました: %w", err)
		return &ogen.APIUsersGetInternalServerError{}, errMsg
	}

	apiUsers := make(ogen.APIUsersGetOKApplicationJSON, len(users))
	for i, user := range users {
		apiUsers[i] = ogen.User{
			ID:        ogen.NewOptInt(int(user.ID)),
			Name:      ogen.NewOptString(user.Name),
			Email:     ogen.NewOptString(user.Email),
			CreatedAt: ogen.NewOptDateTime(user.CreatedAt),
			UpdatedAt: ogen.NewOptDateTime(user.UpdatedAt),
		}
	}

	return &apiUsers, nil
}

func (s *UserService) APIUsersCountGet(ctx context.Context) (ogen.APIUsersCountGetRes, error) {
	var count int64
	if err := db.DB.WithContext(ctx).Model(&models.User{}).Count(&count).Error; err != nil {
		errMsg := fmt.Errorf("データベースからのユーザー数の取得に失敗しました: %w", err)
		return &ogen.APIUsersCountGetInternalServerError{}, errMsg
	}
	return &ogen.CountResponse{Count: ogen.NewOptInt(int(count))}, nil
}
