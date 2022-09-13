package routes

import (
	"example.choreo.dev/internal/config"
	"example.choreo.dev/internal/controllers"
	"example.choreo.dev/internal/repositories"
)

var petController *controllers.PetController
var categoryController *controllers.CategoryController

func initControllers() {
	initialData := config.LoadInitialData()
	categoryRepository := repositories.NewCategoryRepository(initialData.Categories)
	petRepository := repositories.NewPetRepository(initialData.Pets)
	petController = controllers.NewPetController(petRepository, categoryRepository)
	categoryController = controllers.NewCategoryController(categoryRepository)
}
