package models

import (
	"time"

	"gorm.io/gorm"
)

type ModelSchedule struct {
	ID           uint           `json:"id" gorm:"primaryKey"`
	ModelID      uint           `gorm:"index;not null" json:"-"`
	ScheduleDate time.Time      `json:"schedule_date" gorm:"not null"`
	CreatedAt    time.Time      `json:"created_at" gorm:"autoCreateTime"`
	UpdatedAt    time.Time      `json:"updated_at" gorm:"autoUpdateTime"`
	DeletedAt    gorm.DeletedAt `json:"deleted_at" gorm:"index"`

	ModelCatalogues *ModelCatalogue `gorm:"foreignKey:ModelID;constraint:onUpdate:CASCADE,onDelete:CASCADE" json:"model"`
}
