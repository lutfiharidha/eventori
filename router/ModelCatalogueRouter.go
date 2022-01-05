package router

import (
	"eventori/app/controllers"
	"eventori/app/repositories"
	"eventori/app/services"
	"eventori/config"

	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

var (
	db *gorm.DB = config.SetupDatabaseConnection()

	modelCatalogueRepository repositories.ModelCatalogueRepository = repositories.NewModelCatalogueRepository(db)

	modelCatalogueService services.ModelCatalogueService = services.NewModelCatalogueService(modelCatalogueRepository)

	modelCatalogueController controllers.ModelCatalogueController = controllers.NewModelCatalogueController(modelCatalogueService)
)

func ModelCatalogueRoute(route *echo.Echo) {

	modelCatalogueRoutes := route.Group("models")
	{
		modelCatalogueRoutes.GET("/list", modelCatalogueController.All)
		modelCatalogueRoutes.POST("/create", modelCatalogueController.Insert)
		modelCatalogueRoutes.GET("/:id", modelCatalogueController.FindByID)
		modelCatalogueRoutes.PATCH("/:id", modelCatalogueController.Update)
		modelCatalogueRoutes.DELETE("/:id", modelCatalogueController.Delete)
	}
}
