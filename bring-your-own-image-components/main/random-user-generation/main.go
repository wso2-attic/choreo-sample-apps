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

package main

import (
	"fmt"
	"math/rand"
	"time"

	faker "github.com/bxcodec/faker/v4"
)

// User represents a user's information
type User struct {
	Name  string `faker:"name"`
	Email string `faker:"email"`
	Age   int
}

func main() {
	// Seed the random number generator
	rand.Seed(time.Now().UnixNano())

	// Generate random user information
	var user User
	err := faker.FakeData(&user)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	// Generate a random age between 18 and 80
	user.Age = rand.Intn(63) + 18

	// Print the generated user information
	fmt.Println("Generated User Information:")
	fmt.Println("Name:", user.Name)
	fmt.Println("Email:", user.Email)
	fmt.Println("Age:", user.Age)
}
