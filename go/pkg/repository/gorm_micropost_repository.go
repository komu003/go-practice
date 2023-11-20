package repository

import (
	"app/models"
	"context"
	"gorm.io/gorm"
)

type GormMicropostRepository struct {
	db *gorm.DB
}

func NewGormMicropostRepository(db *gorm.DB) *GormMicropostRepository {
	return &GormMicropostRepository{db: db}
}

func (r *GormMicropostRepository) GetMicroposts(ctx context.Context) ([]models.Micropost, error) {
	var microposts []models.Micropost
	if err := r.db.WithContext(ctx).Preload("User").Find(&microposts).Error; err != nil {
		return nil, err
	}
	return microposts, nil
}

func (r *GormMicropostRepository) CountMicroposts(ctx context.Context) (int64, error) {
	var count int64
	if err := r.db.WithContext(ctx).Model(&models.Micropost{}).Count(&count).Error; err != nil {
		return 0, err
	}
	return count, nil
}
