package repositories

import (
	// "github.com/ydhnwb/golang_api/entity"

	"eventori/app/models"

	"gorm.io/gorm"
)

type ModelScheduleRepository interface {
	InsertSchedule(wh models.ModelSchedule) models.ModelSchedule
	UpdateSchedule(wh models.ModelSchedule) models.ModelSchedule
	DeleteSchedule(wh models.ModelSchedule) models.ModelSchedule
	AllSchedule() []models.ModelSchedule
	FindScheduleByModel(modelID string) models.ModelSchedule
}

type modelScheduleConnection struct {
	connection *gorm.DB
}

func NewModelScheduleRepository(dbConn *gorm.DB) ModelScheduleRepository {
	return &modelScheduleConnection{
		connection: dbConn,
	}
}

func (db *modelScheduleConnection) InsertSchedule(model models.ModelSchedule) models.ModelSchedule {
	db.connection.Create(&model)
	db.connection.Preload("Warehouse").Find(&model)
	return model
}

func (db *modelScheduleConnection) UpdateSchedule(model models.ModelSchedule) models.ModelSchedule {
	db.connection.Save(&model)
	db.connection.Preload("Warehouse").Find(&model)
	return model
}

func (db *modelScheduleConnection) DeleteSchedule(model models.ModelSchedule) models.ModelSchedule {
	db.connection.Delete(&model)
	db.connection.Preload("Warehouse").Find(&model)
	return model
}

func (db *modelScheduleConnection) AllSchedule() []models.ModelSchedule {
	var models []models.ModelSchedule
	db.connection.Preload("Model").Find(&models)
	return models
}

func (db *modelScheduleConnection) FindScheduleByModel(modelID string) models.ModelSchedule {
	var model models.ModelSchedule
	db.connection.Where("model_id =?", modelID).Preload("Model").First(&model)
	return model
}
