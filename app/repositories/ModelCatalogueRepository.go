package repositories

import (
	// "github.com/ydhnwb/golang_api/entity"

	"eventori/app/models"

	"gorm.io/gorm"
)

type ModelCatalogueRepository interface {
	InsertModel(wh models.ModelCatalogue) models.ModelCatalogue
	UpdateModel(wh models.ModelCatalogue) models.ModelCatalogue
	DeleteModel(wh models.ModelCatalogue) models.ModelCatalogue
	AllModel() []models.ModelCatalogue
	FindModelByID(modelID string) models.ModelCatalogue
}

type modelCatalogueConnection struct {
	connection *gorm.DB
}

func NewModelCatalogueRepository(dbConn *gorm.DB) ModelCatalogueRepository {
	return &modelCatalogueConnection{
		connection: dbConn,
	}
}

func (db *modelCatalogueConnection) InsertModel(model models.ModelCatalogue) models.ModelCatalogue {
	db.connection.Create(&model)
	db.connection.Preload("Warehouse").Find(&model)
	return model
}

func (db *modelCatalogueConnection) UpdateModel(model models.ModelCatalogue) models.ModelCatalogue {
	db.connection.Save(&model)
	db.connection.Preload("Warehouse").Find(&model)
	return model
}

func (db *modelCatalogueConnection) DeleteModel(model models.ModelCatalogue) models.ModelCatalogue {
	db.connection.Delete(&model)
	db.connection.Preload("Warehouse").Find(&model)
	return model
}

func (db *modelCatalogueConnection) AllModel() []models.ModelCatalogue {
	var models []models.ModelCatalogue
	db.connection.Find(&models)
	return models
}

func (db *modelCatalogueConnection) FindModelByID(modelID string) models.ModelCatalogue {
	var model models.ModelCatalogue
	db.connection.Where("id =?", modelID).First(&model)
	return model
}
