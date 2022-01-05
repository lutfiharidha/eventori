package database

import (
	"eventori/app/models"
	"math/rand"
	"time"
)

func ModelScheduleSeeder() {
	for i := 0; i < 10; i++ {
		db.Create(&models.ModelSchedule{
			ModelID:      uint(rand.Intn(10 - 1)),
			ScheduleDate: time.Now().Add(1000),
		})
	}
}
