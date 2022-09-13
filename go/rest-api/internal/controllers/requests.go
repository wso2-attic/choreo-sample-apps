package controllers

import (
	"example.choreo.dev/internal/models"
)

type AddCategoryRequest struct {
	Name string `json:"name"`
}

type AddCategoryResponse = models.Category

type ListCategoriesResponse = []models.Category

type AddPetRequest struct {
	Name       string `json:"name"`
	CategoryId string `json:"category_id"`
	Available  bool   `json:"available"`
}

type AddPetResponse = models.Pet

type UpdatePetRequest struct {
	Name       *string `json:"name"`
	CategoryId *string `json:"category_id"`
	Available  *bool   `json:"available"`
}

type UpdatePetResponse = models.Pet

type GetPetByIdResponse = models.Pet
