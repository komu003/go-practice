package repository

import (
	"app/models"
	"context"
)

type UserRepository interface {
	GetUsers(ctx context.Context) ([]models.User, error)
	CountUsers(ctx context.Context) (int64, error)
}
