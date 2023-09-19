/*
 * Copyright (c) 2023, WSO2 LLC. (https://www.wso2.com/) All Rights Reserved.
 *
 * WSO2 LLC. licenses this file to you under the Apache License,
 * Version 2.0 (the "License"); you may not use this file except
 * in compliance with the License.
 * You may obtain a copy of the License at
 *
 *     http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing,
 * software distributed under the License is distributed on an
 * "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY
 * KIND, either express or implied. See the License for the
 * specific language governing permissions and limitations
 * under the License.
 */

package controllers

import (
	"context"
	"errors"
	"fmt"
	"net/http"

	"github.com/gofiber/fiber/v2"
	"github.com/wso2/choreo-sample-apps/byoi-components/services/pet-store/internal/models"
	"github.com/wso2/choreo-sample-apps/byoi-components/services/pet-store/internal/repositories"
)

type PetController struct {
	petRepository models.PetRepository
}

func NewPetController(petRepository models.PetRepository) *PetController {
	return &PetController{petRepository}
}

func (c *PetController) AddPet(ctx context.Context, newPet models.Pet) (models.Pet, error) {
	if err := validatePet(newPet); err != nil {
		return models.Pet{}, err
	}
	pet, err := c.petRepository.Add(ctx, newPet)
	if errors.Is(err, repositories.ErrRecordAlreadyExists) {
		return models.Pet{}, makeHttpConflictError(newPet.Id)
	} else if err != nil {
		return models.Pet{}, makeHttpInternalServerError()
	}
	return pet, nil
}

func (c *PetController) UpdatePet(ctx context.Context, updatedPet models.Pet) (models.Pet, error) {
	if err := validatePet(updatedPet); err != nil {
		return models.Pet{}, err
	}
	pet, err := c.petRepository.Update(ctx, updatedPet)
	if errors.Is(err, repositories.ErrRecordNotFound) {
		return models.Pet{}, makeHttpNotFoundError(updatedPet.Id)
	} else if err != nil {
		return models.Pet{}, makeHttpInternalServerError()
	}
	return pet, nil
}

func (c *PetController) ListPets(ctx context.Context) ([]models.Pet, error) {
	pets, err := c.petRepository.List(ctx)
	if err != nil {
		return nil, makeHttpInternalServerError()
	}
	if pets == nil {
		return make([]models.Pet, 0), nil
	}
	return pets, nil
}

func (c *PetController) GetPet(ctx context.Context, petId string) (models.Pet, error) {
	pet, err := c.petRepository.GetById(ctx, petId)
	if errors.Is(err, repositories.ErrRecordNotFound) {
		return models.Pet{}, makeHttpNotFoundError(petId)
	} else if err != nil {
		return models.Pet{}, makeHttpInternalServerError()
	}
	return pet, nil
}

func (c *PetController) DeletePet(ctx context.Context, petId string) (models.Pet, error) {
	pet, err := c.petRepository.DeleteById(ctx, petId)
	if errors.Is(err, repositories.ErrRecordNotFound) {
		return models.Pet{}, makeHttpNotFoundError(petId)
	} else if err != nil {
		return models.Pet{}, makeHttpInternalServerError()
	}
	return pet, nil
}

func makeHttpNotFoundError(id string) *fiber.Error {
	return fiber.NewError(http.StatusNotFound, fmt.Sprintf("the pet id [%s] is not found", id))
}

func makeHttpConflictError(id string) *fiber.Error {
	return fiber.NewError(http.StatusConflict, fmt.Sprintf("the pet id [%s] is already exists", id))
}

func makeHttpInternalServerError() *fiber.Error {
	return fiber.NewError(http.StatusInternalServerError, "internal server error")
}

func validatePet(pet models.Pet) *fiber.Error {
	if pet.Name == "" {
		return fiber.NewError(http.StatusBadRequest, "pet title is required")
	}
	return nil
}
