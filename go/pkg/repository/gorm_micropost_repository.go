package repository

import (
	"app/models"
	"app/pkg/db"
	"context"
)

type GormMicropostRepository struct{}

func NewGormMicropostRepository() *GormMicropostRepository {
	return &GormMicropostRepository{}
}

func (r *GormMicropostRepository) GetMicroposts(ctx context.Context) ([]models.Micropost, error) {
	var microposts []models.Micropost
	if err := db.DB.WithContext(ctx).Preload("User").Find(&microposts).Error; err != nil {
		return nil, err
	}
	return microposts, nil
}

func (r *GormMicropostRepository) CountMicroposts(ctx context.Context) (int64, error) {
	var count int64
	if err := db.DB.WithContext(ctx).Model(&models.Micropost{}).Count(&count).Error; err != nil {
		return 0, err
	}
	return count, nil
}
