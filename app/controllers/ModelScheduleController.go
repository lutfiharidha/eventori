package controllers

import (
	"eventori/app/dtos"
	"eventori/app/helpers"
	"eventori/app/models"
	"eventori/app/services"
	"net/http"

	"github.com/labstack/echo/v4"
)

type ModelScheduleController interface {
	All(context echo.Context) error
	FindByID(context echo.Context) error
	Insert(context echo.Context) error
	Update(context echo.Context) error
	Delete(context echo.Context) error
}

type modelScheduleController struct {
	modelScheduleService services.ModelScheduleService
}

func NewModelScheduleController(modelScheduleServ services.ModelScheduleService) ModelScheduleController {
	return &modelScheduleController{
		modelScheduleService: modelScheduleServ,
	}
}

func (c *modelScheduleController) All(context echo.Context) error {
	var modelSchedules []models.ModelSchedule = c.modelScheduleService.All()
	res := helpers.BuildResponse(true, "OK", modelSchedules)
	return context.JSON(http.StatusOK, res)
}

func (c *modelScheduleController) FindByID(context echo.Context) error {
	id := context.Param("id")
	var modelSchedule models.ModelSchedule = c.modelScheduleService.FindByModel(id)
	if (modelSchedule == models.ModelSchedule{}) {
		res := helpers.BuildErrorResponse("Data not found", "No data with given id", helpers.EmptyObj{})
		return context.JSON(http.StatusNotFound, res)
	} else {
		res := helpers.BuildResponse(true, "OK", modelSchedule)
		return context.JSON(http.StatusOK, res)
	}
}

func (c *modelScheduleController) Insert(context echo.Context) error {
	var modelScheduleCreateDTO dtos.ModelScheduleCreateDTO
	context.Bind(&modelScheduleCreateDTO)

	// if errDTO != nil {
	// 	res := helpers.BuildErrorResponse("Failed to process request", errDTO.Error(), helpers.EmptyObj{})
	// 	context.JSON(http.StatusBadRequest, res)
	// }
	result := c.modelScheduleService.Insert(modelScheduleCreateDTO)
	response := helpers.BuildResponse(true, "OK", result)
	return context.JSON(http.StatusCreated, response)
}

func (c *modelScheduleController) Update(context echo.Context) error {
	var modelScheduleUpdateDTO dtos.ModelScheduleUpdateDTO
	context.Bind(&modelScheduleUpdateDTO)
	id := context.Param("id")
	var modelSchedule models.ModelSchedule = c.modelScheduleService.FindByModel(id)
	if (modelSchedule == models.ModelSchedule{}) {
		res := helpers.BuildErrorResponse("Data not found", "No data with given id", helpers.EmptyObj{})
		return context.JSON(http.StatusNotFound, res)
	} else {
		result := c.modelScheduleService.Update(modelScheduleUpdateDTO)
		response := helpers.BuildResponse(true, "OK", result)
		return context.JSON(http.StatusOK, response)
	}
}

func (c *modelScheduleController) Delete(context echo.Context) error {
	var modelSchedule models.ModelSchedule
	id := context.Param("id")
	modelSchedule = c.modelScheduleService.FindByModel(id)
	if (modelSchedule == models.ModelSchedule{}) {
		res := helpers.BuildErrorResponse("Data not found", "No data with given id", helpers.EmptyObj{})
		return context.JSON(http.StatusNotFound, res)
	} else {
		result := c.modelScheduleService.Delete(modelSchedule)
		res := helpers.BuildResponse(true, "Deleted", result)
		return context.JSON(http.StatusOK, res)
	}
}
