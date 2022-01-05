package services

import (
	"eventori/app/dtos"
	"eventori/app/models"
	"eventori/app/repositories"
	"log"

	"github.com/mashingan/smapping"
)

type ModelCatalogueService interface {
	Insert(mod dtos.ModelCatalogueCreateDTO) models.ModelCatalogue
	Update(mod dtos.ModelCatalogueUpdateDTO) models.ModelCatalogue
	Delete(mod models.ModelCatalogue) models.ModelCatalogue
	All() []models.ModelCatalogue
	FindByID(modelID string) models.ModelCatalogue
}

type modelCatalogueService struct {
	modelCatalogueRepository repositories.ModelCatalogueRepository
}

func NewModelCatalogueService(modelCatalogueRepo repositories.ModelCatalogueRepository) ModelCatalogueService {
	return &modelCatalogueService{
		modelCatalogueRepository: modelCatalogueRepo,
	}
}

func (service *modelCatalogueService) Insert(mod dtos.ModelCatalogueCreateDTO) models.ModelCatalogue {
	model := models.ModelCatalogue{}
	err := smapping.FillStruct(&model, smapping.MapFields(&mod))
	if err != nil {
		log.Fatalf("Failed map %v: ", err)
	}
	res := service.modelCatalogueRepository.InsertModel(model)
	return res
}

func (service *modelCatalogueService) Update(mod dtos.ModelCatalogueUpdateDTO) models.ModelCatalogue {
	model := models.ModelCatalogue{}
	err := smapping.FillStruct(&model, smapping.MapFields(&mod))
	if err != nil {
		log.Fatalf("Failed map %v: ", err)
	}
	res := service.modelCatalogueRepository.UpdateModel(model)
	return res
}

func (service *modelCatalogueService) Delete(mod models.ModelCatalogue) models.ModelCatalogue {
	return service.modelCatalogueRepository.DeleteModel(mod)
}

func (service *modelCatalogueService) All() []models.ModelCatalogue {
	return service.modelCatalogueRepository.AllModel()
}

func (service *modelCatalogueService) FindByID(modelID string) models.ModelCatalogue {
	return service.modelCatalogueRepository.FindModelByID(modelID)
}
