package database

import (
	"eventori/app/models"
	"time"
)

func ModelScheduleSeeder() {

	modelCatalogue := models.ModelCatalogue{}
	db.Select("id").Last(&modelCatalogue)

	db.Create(&models.ModelSchedule{
		ModelID:      modelCatalogue.ID,
		ScheduleDate: time.Now().Add(1000),
	})

}
