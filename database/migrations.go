package database

import (
	"eventori/app/models"
	"eventori/config"

	"gorm.io/gorm"
)

var (
	db *gorm.DB = config.SetupDatabaseConnection()
)

func InitialMigration() {
	db.Migrator().DropTable(&models.ModelCatalogue{})
	db.Migrator().CreateTable(&models.ModelCatalogue{})

	db.Migrator().DropTable(&models.ModelSchedule{})
	db.Migrator().CreateTable(&models.ModelSchedule{})
}
