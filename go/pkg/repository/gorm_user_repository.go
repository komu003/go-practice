package repository

import (
	"app/models"
	"app/pkg/db"
	"context"
)

type GormUserRepository struct{}

func NewGormUserRepository() *GormUserRepository {
	return &GormUserRepository{}
}

func (r *GormUserRepository) GetUsers(ctx context.Context) ([]models.User, error) {
	var users []models.User
	if err := db.DB.WithContext(ctx).Find(&users).Error; err != nil {
		return nil, err
	}
	return users, nil
}

func (r *GormUserRepository) CountUsers(ctx context.Context) (int64, error) {
	var count int64
	if err := db.DB.WithContext(ctx).Model(&models.User{}).Count(&count).Error; err != nil {
		return 0, err
	}
	return count, nil
}
