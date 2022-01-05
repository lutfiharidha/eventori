package services

import (
	"eventori/app/dtos"
	"eventori/app/models"
	"eventori/app/repositories"
	"log"

	"github.com/mashingan/smapping"
)

type ModelScheduleService interface {
	Insert(mod dtos.ModelScheduleCreateDTO) models.ModelSchedule
	Update(mod dtos.ModelScheduleUpdateDTO) models.ModelSchedule
	Delete(mod models.ModelSchedule) models.ModelSchedule
	All() []models.ModelSchedule
	FindByModel(modelID string) models.ModelSchedule
}

type modelScheduleService struct {
	modelScheduleRepository repositories.ModelScheduleRepository
}

func NewModelScheduleService(modelScheduleRepo repositories.ModelScheduleRepository) ModelScheduleService {
	return &modelScheduleService{
		modelScheduleRepository: modelScheduleRepo,
	}
}

func (service *modelScheduleService) Insert(mod dtos.ModelScheduleCreateDTO) models.ModelSchedule {
	model := models.ModelSchedule{}
	err := smapping.FillStruct(&model, smapping.MapFields(&mod))
	if err != nil {
		log.Fatalf("Failed map %v: ", err)
	}
	res := service.modelScheduleRepository.InsertSchedule(model)
	return res
}

func (service *modelScheduleService) Update(mod dtos.ModelScheduleUpdateDTO) models.ModelSchedule {
	model := models.ModelSchedule{}
	err := smapping.FillStruct(&model, smapping.MapFields(&mod))
	if err != nil {
		log.Fatalf("Failed map %v: ", err)
	}
	res := service.modelScheduleRepository.UpdateSchedule(model)
	return res
}

func (service *modelScheduleService) Delete(mod models.ModelSchedule) models.ModelSchedule {
	return service.modelScheduleRepository.DeleteSchedule(mod)
}

func (service *modelScheduleService) All() []models.ModelSchedule {
	return service.modelScheduleRepository.AllSchedule()
}

func (service *modelScheduleService) FindByModel(modelID string) models.ModelSchedule {
	return service.modelScheduleRepository.FindScheduleByModel(modelID)
}
