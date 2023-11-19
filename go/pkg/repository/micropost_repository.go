package repository

import (
	"app/models"
	"context"
)

type MicropostRepository interface {
	GetMicroposts(ctx context.Context) ([]models.Micropost, error)
	CountMicroposts(ctx context.Context) (int64, error)
}
