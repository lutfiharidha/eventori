package models

import (
	"time"

	"gorm.io/gorm"
)

type ModelCatalogue struct {
	ID          uint           `json:"id" gorm:"primaryKey"`
	ModelName   string         `json:"model_name" gorm:"not null"`
	ImageUrl    string         `json:"image_url" gorm:"not null"`
	Description string         `json:"description" gorm:"not null"`
	CreatedAt   time.Time      `json:"created_at" gorm:"autoCreateTime"`
	UpdatedAt   time.Time      `json:"updated_at" gorm:"autoUpdateTime"`
	DeletedAt   gorm.DeletedAt `json:"deleted_at" gorm:"index"`
}
