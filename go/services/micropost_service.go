package services

import (
	"app/models"
	"app/ogen"
	"app/pkg/repository"
	"context"
	"fmt"
)

type MicropostService struct {
	repo repository.MicropostRepository
}

func NewMicropostService(repo repository.MicropostRepository) *MicropostService {
	return &MicropostService{
		repo: repo,
	}
}

func (s *MicropostService) APIMicropostsGet(ctx context.Context) (ogen.APIMicropostsGetRes, error) {
	microposts, err := s.repo.GetMicroposts(ctx)
	if err != nil {
		errMsg := fmt.Errorf("データベースからのマイクロポスト一覧の取得に失敗しました: %w", err)
		return &ogen.APIMicropostsGetInternalServerError{}, errMsg
	}

	apiMicroposts := make(ogen.APIMicropostsGetOKApplicationJSON, len(microposts))
	for i, mp := range microposts {
		apiMicroposts[i] = ogen.Micropost{
			ID:        ogen.NewOptInt(int(mp.ID)),
			Content:   ogen.NewOptString(mp.Content),
			UserId:    ogen.NewOptInt(int(mp.UserID)),
			CreatedAt: ogen.NewOptDateTime(mp.CreatedAt),
			UpdatedAt: ogen.NewOptDateTime(mp.UpdatedAt),
			User:      convertUserToOptUser(mp.User),
		}
	}

	return &apiMicroposts, nil
}

func (s *MicropostService) APIMicropostsCountGet(ctx context.Context) (ogen.APIMicropostsCountGetRes, error) {
	count, err := s.repo.CountMicroposts(ctx) // リポジトリを使用してカウントを取得
	if err != nil {
		errMsg := fmt.Errorf("データベースからのマイクロポスト数の取得に失敗しました: %w", err)
		return &ogen.APIMicropostsCountGetInternalServerError{}, errMsg
	}
	return &ogen.CountResponse{Count: ogen.NewOptInt(int(count))}, nil
}

func convertUserToOptUser(user models.User) ogen.OptUser {
	return ogen.OptUser{
		Set: true,
		Value: ogen.User{
			ID:        ogen.NewOptInt(int(user.ID)),
			Name:      ogen.NewOptString(user.Name),
			Email:     ogen.NewOptString(user.Email),
			CreatedAt: ogen.NewOptDateTime(user.CreatedAt),
			UpdatedAt: ogen.NewOptDateTime(user.UpdatedAt),
		},
	}
}
