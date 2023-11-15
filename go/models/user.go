package models

import (
	"time"
)

type User struct {
	ID         uint        `gorm:"primaryKey;autoIncrement"`
	Name       string      `gorm:"type:varchar(255);not null"`
	Email      string      `gorm:"type:varchar(255);uniqueIndex;not null"`
	CreatedAt  time.Time   `gorm:"autoCreateTime"`
	UpdatedAt  time.Time   `gorm:"autoUpdateTime"`
	Microposts []Micropost `gorm:"foreignKey:UserID"`
}
