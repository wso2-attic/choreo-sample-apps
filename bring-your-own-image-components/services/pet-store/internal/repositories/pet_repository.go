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
	"fmt"
	"sync"

	"github.com/google/uuid"

	"github.com/wso2/choreo-sample-apps/bring-your-image-components/services/pet-store/internal/models"
)

type petRepository struct {
	store map[string]models.Pet
	lock  sync.RWMutex
}

func NewPetRepository(initialData []models.Pet) models.PetRepository {
	m := make(map[string]models.Pet, 0)
	if len(initialData) > 0 {
		for _, pet := range initialData {
			m[pet.Id] = pet
		}
	}
	return &petRepository{
		store: m,
		lock:  sync.RWMutex{},
	}
}

func (r *petRepository) Add(ctx context.Context, pet models.Pet) (models.Pet, error) {
	r.lock.Lock()
	defer r.lock.Unlock()
	if pet.Id == "" {
		pet.Id = uuid.NewString()
	}
	if _, ok := r.store[pet.Id]; ok {
		return models.Pet{}, fmt.Errorf("petRepository:Add: %w", ErrRecordAlreadyExists)
	}
	r.store[pet.Id] = pet
	return r.store[pet.Id], nil
}

func (r *petRepository) Update(ctx context.Context, updatedPet models.Pet) (models.Pet, error) {
	r.lock.Lock()
	defer r.lock.Unlock()
	if _, ok := r.store[updatedPet.Id]; !ok {
		return models.Pet{}, fmt.Errorf("petRepository:Update: %w", ErrRecordNotFound)
	}
	r.store[updatedPet.Id] = updatedPet
	return r.store[updatedPet.Id], nil
}

func (r *petRepository) List(ctx context.Context) ([]models.Pet, error) {
	r.lock.RLock()
	defer r.lock.RUnlock()
	var pets []models.Pet
	for _, pet := range r.store {
		pets = append(pets, pet)
	}
	return pets, nil
}

func (r *petRepository) GetById(ctx context.Context, id string) (models.Pet, error) {
	r.lock.RLock()
	defer r.lock.RUnlock()
	if _, ok := r.store[id]; !ok {
		return models.Pet{}, fmt.Errorf("petRepository:GetById: %w", ErrRecordNotFound)
	}
	return r.store[id], nil
}

func (r *petRepository) DeleteById(ctx context.Context, id string) (models.Pet, error) {
	r.lock.RLock()
	defer r.lock.RUnlock()

	if _, ok := r.store[id]; !ok {
		return models.Pet{}, fmt.Errorf("petRepository:DeleteById: %w", ErrRecordNotFound)
	}
	pet := r.store[id]
	delete(r.store, id)
	return pet, nil
}
