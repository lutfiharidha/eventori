package dtos

import "time"

type ModelScheduleCreateDTO struct {
	ID           int64     `json:"id" form:"id"`
	ModelID      int64     `json:"model_id" form:"model_id" binding:"required"`
	ScheduleDate time.Time `json:"schadule_date" form:"schadule_date" binding:"required"`
}

type ModelScheduleUpdateDTO struct {
	ID           int64     `json:"id" form:"id"`
	ModelName    int64     `json:"model_id" form:"model_id" binding:"required"`
	ScheduleDate time.Time `json:"schadule_date" form:"schadule_date" binding:"required"`
}

type ModelScheduleRestoreDTO struct {
	ID int64 `json:"id" form:"id"`
}
