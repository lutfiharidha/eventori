package router

import (
	"eventori/app/controllers"
	"eventori/app/repositories"
	"eventori/app/services"

	"github.com/labstack/echo/v4"
)

var (
	modelScheduleRepository repositories.ModelScheduleRepository = repositories.NewModelScheduleRepository(db)

	modelScheduleService services.ModelScheduleService = services.NewModelScheduleService(modelScheduleRepository)

	modelScheduleController controllers.ModelScheduleController = controllers.NewModelScheduleController(modelScheduleService)
)

func ModelScheduleRoute(route *echo.Echo) {

	modelScheduleRoutes := route.Group("models/schedules")
	{
		modelScheduleRoutes.GET("/list", modelScheduleController.All)
		modelScheduleRoutes.POST("/create", modelScheduleController.Insert)
		modelScheduleRoutes.GET("/:model_id", modelScheduleController.FindByModel)
		modelScheduleRoutes.PATCH("/:id", modelScheduleController.Update)
		modelScheduleRoutes.DELETE("/:id", modelScheduleController.Delete)
	}
}
