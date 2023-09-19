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

package models

import (
	"context"
)

type Pet struct {
	Id   string `json:"id" example:"fe2594d0-ccea-42a2-97ac-0487458b5642"`
	Name string `json:"name" example:"Kitty"`
	Age  int    `json:"age" example:"8"`
}

type PetRepository interface {
	Add(ctx context.Context, pet Pet) (Pet, error)
	Update(ctx context.Context, updatedPet Pet) (Pet, error)
	List(ctx context.Context) ([]Pet, error)
	GetById(ctx context.Context, id string) (Pet, error)
	DeleteById(ctx context.Context, id string) (Pet, error)
}
