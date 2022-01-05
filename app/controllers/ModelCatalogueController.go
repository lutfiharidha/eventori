package controllers

import (
	"eventori/app/dtos"
	"eventori/app/helpers"
	"eventori/app/models"
	"eventori/app/services"
	"net/http"

	"github.com/labstack/echo/v4"
)

type ModelCatalogueController interface {
	All(context echo.Context) error
	FindByID(context echo.Context) error
	Insert(context echo.Context) error
	Update(context echo.Context) error
	Delete(context echo.Context) error
}

type modelCatalogueController struct {
	modelCatalogueService services.ModelCatalogueService
}

func NewModelCatalogueController(modelCatalogueServ services.ModelCatalogueService) ModelCatalogueController {
	return &modelCatalogueController{
		modelCatalogueService: modelCatalogueServ,
	}
}

func (c modelCatalogueController) All(context echo.Context) error {
	var modelCatalogues []models.ModelCatalogue = c.modelCatalogueService.All()
	res := helpers.BuildResponse(true, "OK", modelCatalogues)
	return context.JSON(http.StatusOK, res)
}

func (c *modelCatalogueController) FindByID(context echo.Context) error {
	id := context.Param("id")
	var modelCatalogue models.ModelCatalogue = c.modelCatalogueService.FindByID(id)
	if (modelCatalogue == models.ModelCatalogue{}) {
		res := helpers.BuildErrorResponse("Data not found", "No data with given id", helpers.EmptyObj{})
		return context.JSON(http.StatusNotFound, res)
	} else {
		res := helpers.BuildResponse(true, "OK", modelCatalogue)
		return context.JSON(http.StatusOK, res)
	}
}

func (c *modelCatalogueController) Insert(context echo.Context) error {
	var modelCatalogueCreateDTO dtos.ModelCatalogueCreateDTO
	context.Bind(&modelCatalogueCreateDTO)

	// if errDTO != nil {
	// 	res := helpers.BuildErrorResponse("Failed to process request", errDTO.Error(), helpers.EmptyObj{})
	// 	context.JSON(http.StatusBadRequest, res)
	// }
	result := c.modelCatalogueService.Insert(modelCatalogueCreateDTO)
	response := helpers.BuildResponse(true, "OK", result)
	return context.JSON(http.StatusCreated, response)
}

func (c *modelCatalogueController) Update(context echo.Context) error {
	var modelCatalogueUpdateDTO dtos.ModelCatalogueUpdateDTO
	context.Bind(&modelCatalogueUpdateDTO)
	id := context.Param("id")
	var modelCatalogue models.ModelCatalogue = c.modelCatalogueService.FindByID(id)
	if (modelCatalogue == models.ModelCatalogue{}) {
		res := helpers.BuildErrorResponse("Data not found", "No data with given id", helpers.EmptyObj{})
		return context.JSON(http.StatusNotFound, res)
	} else {
		result := c.modelCatalogueService.Update(modelCatalogueUpdateDTO)
		response := helpers.BuildResponse(true, "OK", result)
		return context.JSON(http.StatusOK, response)
	}
}

func (c *modelCatalogueController) Delete(context echo.Context) error {
	var modelCatalogue models.ModelCatalogue
	id := context.Param("id")
	modelCatalogue = c.modelCatalogueService.FindByID(id)
	if (modelCatalogue == models.ModelCatalogue{}) {
		res := helpers.BuildErrorResponse("Data not found", "No data with given id", helpers.EmptyObj{})
		return context.JSON(http.StatusNotFound, res)
	} else {
		result := c.modelCatalogueService.Delete(modelCatalogue)
		res := helpers.BuildResponse(true, "Deleted", result)
		return context.JSON(http.StatusOK, res)
	}
}
