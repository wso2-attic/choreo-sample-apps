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
package repositories

import (
	"context"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/wso2/choreo-sample-apps/bring-your-image-components/services/pet-store/internal/models"
)

func TestPetRepository(t *testing.T) {
	// Create a new repository for testing with an initial book.
	initialPet := models.Pet{Id: "1", Name: "Test Pet", Age: 10}
	updatedPet := models.Pet{Id: initialPet.Id, Name: "Updated Pet", Age: 11}
	repo := NewPetRepository([]models.Pet{initialPet})

	t.Run("Add", func(t *testing.T) {
		// Test adding a new book.
		newPet := models.Pet{Name: "Test Pet", Age: 10}
		addedPet, err := repo.Add(context.Background(), newPet)
		assert.NoError(t, err)
		assert.NotEmpty(t, addedPet.Id)

		// Test adding a book with an existing ID.
		duplicatePet := models.Pet{Id: addedPet.Id, Name: "Duplicate Pet", Age: 10}
		_, err = repo.Add(context.Background(), duplicatePet)
		assert.Error(t, err)
	})

	t.Run("Update", func(t *testing.T) {
		// Test updating an existing book.
		updated, err := repo.Update(context.Background(), updatedPet)
		assert.NoError(t, err)
		assert.Equal(t, updatedPet, updated)

		// Test updating a non-existing book.
		nonExistingPet := models.Pet{Id: "non-existing-id", Name: "Non Existing Pet", Age: 10}
		_, err = repo.Update(context.Background(), nonExistingPet)
		assert.Error(t, err)
	})

	t.Run("List", func(t *testing.T) {
		// Test listing all pets.
		pets, err := repo.List(context.Background())
		assert.NoError(t, err)
		assert.Len(t, pets, 2) // Includes the initial book and the one added.
	})

	t.Run("GetById", func(t *testing.T) {
		// Test getting a book by ID.
		book, err := repo.GetById(context.Background(), initialPet.Id)
		assert.NoError(t, err)
		assert.Equal(t, updatedPet, book)

		// Test getting a non-existing book.
		_, err = repo.GetById(context.Background(), "non-existing-id")
		assert.Error(t, err)
	})

	t.Run("DeleteById", func(t *testing.T) {
		// Test deleting a book by ID.
		deletedPet, err := repo.DeleteById(context.Background(), initialPet.Id)
		assert.NoError(t, err)
		assert.Equal(t, updatedPet, deletedPet)

		// Test deleting a non-existing book.
		_, err = repo.DeleteById(context.Background(), "non-existing-id")
		assert.Error(t, err)
	})
}
