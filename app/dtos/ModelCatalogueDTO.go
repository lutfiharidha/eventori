package dtos

type ModelCatalogueCreateDTO struct {
	ID          int64  `json:"id" form:"id"`
	ModelName   string `json:"model_name" form:"model_name" binding:"required"`
	ImageUrl    string `json:"image_url" form:"image_url" binding:"required"`
	Description string `json:"description" form:"description"`
}

type ModelCatalogueUpdateDTO struct {
	ID          int64  `json:"id" form:"id"`
	ModelName   string `json:"model_name" form:"model_name" binding:"required"`
	ImageUrl    string `json:"image_url" form:"image_url" binding:"required"`
	Description string `json:"description" form:"description"`
}

type ModelCatalogueRestoreDTO struct {
	ID int64 `json:"id" form:"id"`
}
