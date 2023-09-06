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

package routes

import (
	"fmt"
	"net/http"

	"github.com/gofiber/fiber/v2"
	"github.com/wso2/choreo-sample-apps/go/pet-store/internal/models"

	"github.com/wso2/choreo-sample-apps/go/pet-store/internal/utils"
)

func registerPetStoreRoutes(router fiber.Router) {
	r := router.Group("/pets")
	r.Post("/", AddPet)
	r.Get("/:id", GetPet)
	r.Put("/:id", UpdatePet)
	r.Delete("/:id", DeletePet)
	r.Get("/", ListPets)
}

// AddPet
//
//	@Summary	Add a new pet to the pet store
//	@Tags		pets
//	@Accept		json
//	@Produce	json
//	@Param		request	body	models.Pet	true	"New pet details"
//	@Router		/pets [post]
//	@Success	201	{object}	models.Pet			"successful operation"
//	@Failure	400	{object}	utils.ErrorResponse	"invalid pet details"
//	@Failure	409	{object}	utils.ErrorResponse	"pet already exists"
func AddPet(c *fiber.Ctx) error {
	ctx := utils.GetRequestContext(c)
	newPet := models.Pet{}
	if err := c.BodyParser(&newPet); err != nil {
		return makeHttpBadRequestError(err)
	}
	res, err := petController.AddPet(ctx, newPet)
	if err != nil {
		return err
	}
	return c.Status(fiber.StatusCreated).JSON(res)
}

// UpdatePet
//
//	@Summary	Update a pet store pet by id
//	@Tags		pets
//	@Accept		json
//	@Produce	json
//	@Param		id		path	string		true	"Pet ID"
//	@Param		request	body	models.Pet	true	"Updated pet details"
//	@Router		/pets/{id} [put]
//	@Success	200	{object}	models.Pet			"successful operation"
//	@Failure	400	{object}	utils.ErrorResponse	"invalid pet details"
//	@Failure	404	{object}	utils.ErrorResponse	"pet not found"
func UpdatePet(c *fiber.Ctx) error {
	ctx := utils.GetRequestContext(c)
	id := c.Params("id")
	updatedPet := models.Pet{}
	if err := c.BodyParser(&updatedPet); err != nil {
		return makeHttpBadRequestError(err)
	}
	updatedPet.Id = id
	pet, err := petController.UpdatePet(ctx, updatedPet)
	if err != nil {
		return err
	}
	return c.Status(fiber.StatusOK).JSON(pet)
}

// DeletePet
//
//	@Summary	Delete a pet store pet by id
//	@Tags		pets
//	@Produce	json
//	@Param		id	path	string	true	"Pet ID"
//	@Router		/pets/{id} [delete]
//	@Success	200	{object}	models.Pet			"successful operation"
//	@Failure	404	{object}	utils.ErrorResponse	"pet not found"
func DeletePet(c *fiber.Ctx) error {
	ctx := utils.GetRequestContext(c)
	id := c.Params("id")
	pet, err := petController.DeletePet(ctx, id)
	if err != nil {
		return err
	}
	return c.Status(fiber.StatusOK).JSON(pet)
}

// GetPet
//
//	@Summary	Get pet store pet by id
//
//	@Tags		pets
//
//	@Produce	json
//	@Param		id	path	string	true	"Pet ID"
//	@Router		/pets/{id} [get]
//	@Success	200	{object}	models.Pet			"successful operation"
//	@Failure	404	{object}	utils.ErrorResponse	"pet not found"
func GetPet(c *fiber.Ctx) error {
	ctx := utils.GetRequestContext(c)
	id := c.Params("id")

	pet, err := petController.GetPet(ctx, id)
	if err != nil {
		return err
	}
	return c.Status(fiber.StatusOK).JSON(pet)
}

// ListPets
//
//	@Summary	List all the pet store pets
//	@Tags		pets
//	@Produce	json
//	@Router		/pets [get]
//	@Success	200	{array}	models.Pet	"successful operation"
func ListPets(c *fiber.Ctx) error {
	ctx := utils.GetRequestContext(c)
	pets, err := petController.ListPets(ctx)
	if err != nil {
		return err
	}
	return c.Status(fiber.StatusOK).JSON(pets)
}

func makeHttpBadRequestError(err error) *fiber.Error {
	return fiber.NewError(http.StatusBadRequest, fmt.Sprintf("failed to parse the payload: %s", err.Error()))
}
