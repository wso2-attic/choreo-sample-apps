package routes

import (
	"github.com/gofiber/fiber/v2"

	"example.choreo.dev/internal/controllers"
	"example.choreo.dev/internal/utils"
)

func registerCategoryRoutes(router fiber.Router) {
	r := router.Group("/category")
	r.Post("/", CreateCategory)
	r.Get("/", ListCategories)
}

// CreateCategory
// @Summary      Add pet to the store
// @Accept       json
// @Produce      json
// @Param request body controllers.AddCategoryRequest true "pet details"
// @Router /api/v1/category [post]
// @Failure 409  {object}  any
// @Success 200 {object} controllers.AddCategoryResponse "ok"
func CreateCategory(c *fiber.Ctx) error {
	ctx := utils.GetRequestContext(c)
	var req controllers.AddCategoryRequest

	if err := c.BodyParser(&req); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"status":  "error",
			"details": "failed to parse payload",
		})
	}

	res, err := categoryController.AddCategory(ctx, req)
	if err != nil {
		return err
	}
	return c.Status(fiber.StatusOK).JSON(res)
}

// ListCategories
// @Summary      List available categories
// @Produce      json
// @Router /api/v1/category [get]
// @Success 200 {object} controllers.ListCategoriesResponse "ok"
func ListCategories(c *fiber.Ctx) error {
	ctx := utils.GetRequestContext(c)

	res, err := categoryController.ListCategories(ctx)
	if err != nil {
		return err
	}
	return c.Status(fiber.StatusOK).JSON(res)
}
