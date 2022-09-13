package controllers

import (
	"context"
	"errors"
	"fmt"
	"net/http"

	"github.com/gofiber/fiber/v2"
	"github.com/jinzhu/copier"

	"example.choreo.dev/internal/models"
	"example.choreo.dev/internal/repositories"
)

type PetController struct {
	petRepository      models.PetRepository
	categoryRepository models.CategoryRepository
}

func NewPetController(petRepository models.PetRepository, categoryRepository models.CategoryRepository) *PetController {
	return &PetController{petRepository, categoryRepository}
}

func (c PetController) AddPet(ctx context.Context, request AddPetRequest) (*AddPetResponse, error) {
	var pet models.Pet
	if err := copier.Copy(&pet, request); err != nil {
		return nil, err
	}
	category, err := c.categoryRepository.FindById(ctx, request.CategoryId)
	if err != nil {
		return nil, err
	}
	if category == nil {
		return nil, fiber.NewError(http.StatusBadRequest, fmt.Sprintf("invalid category id [%s]", request.CategoryId))
	}
	if err := c.petRepository.Add(ctx, &pet); err != nil {
		return nil, err
	}
	return &pet, nil
}

func (c PetController) GetPetById(ctx context.Context, petId string) (*GetPetByIdResponse, error) {
	pet, err := c.petRepository.GetById(ctx, petId)
	if err != nil {
		if errors.Is(err, repositories.ErrRecordNotFound) {
			return nil, fiber.NewError(http.StatusNotFound, "record not found")
		}
		return nil, err
	}
	return pet, nil
}
