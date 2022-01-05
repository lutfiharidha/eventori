package main

import (
	"eventori/app/helpers"
	"eventori/config"
	"eventori/router"
	"os"

	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

var db *gorm.DB = config.SetupDatabaseConnection()

func main() {
	defer config.CloseDatabaseConnection(db)
	helpers.Command()
	e := echo.New()
	router.ModelCatalogueRoute(e) //Added all warehouse routes
	router.ModelScheduleRoute(e)  //Added all warehouse routes
	e.Logger.Fatal(e.Start(":" + os.Getenv("APP_PORT")))
}
