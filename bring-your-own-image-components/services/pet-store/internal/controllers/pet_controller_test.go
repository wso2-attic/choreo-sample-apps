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
	"net/http"
	"testing"

	"github.com/gofiber/fiber/v2"
	"github.com/stretchr/testify/assert"
	"github.com/wso2/choreo-sample-apps/bring-your-image-components/services/pet-store/internal/models"
	"github.com/wso2/choreo-sample-apps/bring-your-image-components/services/pet-store/internal/repositories"
)

// MockPetRepository is a mock implementation of the models.PetRepository interface for testing purposes.
type MockPetRepository struct {
	data   map[string]models.Pet
	err    error
	exists bool
}

func (m *MockPetRepository) Add(ctx context.Context, pet models.Pet) (models.Pet, error) {
	if m.exists {
		return models.Pet{}, repositories.ErrRecordAlreadyExists
	}
	return pet, m.err
}

func (m *MockPetRepository) Update(ctx context.Context, pet models.Pet) (models.Pet, error) {
	if !m.exists {
		return models.Pet{}, repositories.ErrRecordNotFound
	}
	return pet, m.err
}

func (m *MockPetRepository) List(ctx context.Context) ([]models.Pet, error) {
	pets := make([]models.Pet, 0, len(m.data))
	for _, pet := range m.data {
		pets = append(pets, pet)
	}
	return pets, m.err
}

func (m *MockPetRepository) GetById(ctx context.Context, id string) (models.Pet, error) {
	pet, ok := m.data[id]
	if !ok {
		return models.Pet{}, repositories.ErrRecordNotFound
	}
	return pet, m.err
}

func (m *MockPetRepository) DeleteById(ctx context.Context, id string) (models.Pet, error) {
	pet, ok := m.data[id]
	if !ok {
		return models.Pet{}, repositories.ErrRecordNotFound
	}
	delete(m.data, id)
	return pet, m.err
}

func TestPetController(t *testing.T) {
	// Create a mock repository for testing.
	mockRepo := &MockPetRepository{
		data:   make(map[string]models.Pet),
		err:    nil,
		exists: false,
	}

	controller := NewPetController(mockRepo)

	t.Run("AddPet", func(t *testing.T) {
		// Test adding a new pet.
		newPet := models.Pet{Name: "New Pet", Age: 10}
		mockRepo.exists = false
		pet, err := controller.AddPet(context.Background(), newPet)
		assert.NoError(t, err)
		assert.Equal(t, newPet.Name, pet.Name)

		// Test adding a pet that already exists.
		mockRepo.exists = true
		_, err = controller.AddPet(context.Background(), newPet)
		assert.Equal(t, fiber.NewError(http.StatusConflict, "the pet id [] is already exists"), err)
	})

	t.Run("UpdatePet", func(t *testing.T) {
		// Test updating an existing pet.
		updatedPet := models.Pet{Id: "1", Name: "Updated Pet", Age: 10}
		mockRepo.exists = true
		pet, err := controller.UpdatePet(context.Background(), updatedPet)
		assert.NoError(t, err)
		assert.Equal(t, updatedPet.Name, pet.Name)

		// Test updating a pet that does not exist.
		mockRepo.exists = false
		_, err = controller.UpdatePet(context.Background(), updatedPet)
		assert.Equal(t, fiber.NewError(http.StatusNotFound, "the pet id [1] is not found"), err)
	})

	t.Run("ListPets", func(t *testing.T) {
		// Test listing pets.
		mockRepo.data = map[string]models.Pet{
			"1": {Id: "1", Name: "Pet 1", Age: 10},
			"2": {Id: "2", Name: "Pet 2", Age: 12},
		}
		pets, err := controller.ListPets(context.Background())
		assert.NoError(t, err)
		assert.Len(t, pets, 2)

		// Test listing pets with an error.
		mockRepo.err = errors.New("mock error")
		_, err = controller.ListPets(context.Background())
		assert.Equal(t, fiber.NewError(http.StatusInternalServerError, "internal server error"), err)
	})

	t.Run("GetPet", func(t *testing.T) {
		// Test getting an existing pet.
		mockRepo.data = map[string]models.Pet{"1": {Id: "1", Name: "Pet 1", Age: 12}}
		mockRepo.err = nil
		pet, err := controller.GetPet(context.Background(), "1")
		assert.NoError(t, err)
		assert.Equal(t, "Pet 1", pet.Name)

		// Test getting a pet that does not exist.
		_, err = controller.GetPet(context.Background(), "2")
		assert.Equal(t, fiber.NewError(http.StatusNotFound, "the pet id [2] is not found"), err)
	})

	t.Run("DeletePet", func(t *testing.T) {
		// Test deleting an existing pet.
		mockRepo.data = map[string]models.Pet{"1": {Id: "1", Name: "Pet 1", Age: 10}}
		pet, err := controller.DeletePet(context.Background(), "1")
		assert.NoError(t, err)
		assert.Equal(t, "Pet 1", pet.Name)

		// Test deleting a pet that does not exist.
		_, err = controller.DeletePet(context.Background(), "2")
		assert.Equal(t, fiber.NewError(http.StatusNotFound, "the pet id [2] is not found"), err)
	})
}
