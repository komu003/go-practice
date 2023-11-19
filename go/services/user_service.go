package services

import (
	"app/ogen"
	"app/pkg/repository"
	"context"
	"fmt"
)

type UserService struct {
	repo repository.UserRepository
}

func NewUserService(repo repository.UserRepository) *UserService {
	return &UserService{
		repo: repo,
	}
}

func (s *UserService) APIUsersGet(ctx context.Context) (ogen.APIUsersGetRes, error) {
	users, err := s.repo.GetUsers(ctx)
	if err != nil {
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
	count, err := s.repo.CountUsers(ctx)
	if err != nil {
		errMsg := fmt.Errorf("データベースからのユーザー数の取得に失敗しました: %w", err)
		return &ogen.APIUsersCountGetInternalServerError{}, errMsg
	}
	return &ogen.CountResponse{Count: ogen.NewOptInt(int(count))}, nil
}
