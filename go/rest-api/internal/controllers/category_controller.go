package controllers

import (
	"context"
	"fmt"
	"net/http"

	"github.com/gofiber/fiber/v2"
	"github.com/jinzhu/copier"

	"example.choreo.dev/internal/models"
)

type CategoryController struct {
	categoryRepository models.CategoryRepository
}

func NewCategoryController(categoryRepository models.CategoryRepository) *CategoryController {
	return &CategoryController{categoryRepository: categoryRepository}
}

func (c CategoryController) AddCategory(ctx context.Context, request AddCategoryRequest) (*AddCategoryResponse, error) {
	var category models.Category
	if err := copier.Copy(&category, request); err != nil {
		return nil, err
	}
	existing, err := c.categoryRepository.FindOneByName(ctx, request.Name)
	if err != nil {
		return nil, err
	}
	if existing != nil {
		return nil, fiber.NewError(http.StatusConflict, fmt.Sprintf("category [%s] already exists", request.Name))
	}
	if err := c.categoryRepository.Add(ctx, &category); err != nil {
		return nil, err
	}
	return &category, nil
}

func (c CategoryController) ListCategories(ctx context.Context) (ListCategoriesResponse, error) {
	return c.categoryRepository.GetAll(ctx)
}
