package route

import (
	"ncbs/api/handler"
	"ncbs/api/model"
	"net/http"

	"github.com/labstack/echo/v4"
)

func SetUpRoutes(e *echo.Echo, recipeHandler *handler.RecipeHandler) {

	e.GET("/", func(c echo.Context) error {
		return c.HTML(http.StatusOK, "Nates chicken bullshit")
	})

	e.GET("/health", func(c echo.Context) error {
		return c.JSON(http.StatusOK, struct{ Status string }{Status: "OK"})
	})

	// could probably have some sort of default middleware that handles the
	// model -> json conversion, along with standard error handling?
	e.GET("/recipes", func(c echo.Context) error {
		recipes, err := recipeHandler.GetAllRecipes()
		if err != nil {
			return c.JSON(http.StatusInternalServerError, map[string]string{
				"error": "Failed to get recipes",
			})
		}
		return c.JSON(http.StatusOK, recipes)
	})

	e.POST("/recipes", func(c echo.Context) error {
		recipe := new(model.Recipe)
		if err := c.Bind(recipe); err != nil {
			return c.JSON(http.StatusBadRequest, map[string]string{
				"error": "Failed to bind request body to Recipe struct",
			})
		}

		err := recipeHandler.CreateRecipe(*recipe)
		if err != nil {
			return c.JSON(http.StatusInternalServerError, map[string]string{
				"error": "Failed to create recipe",
			})
		}

		return c.JSON(http.StatusOK, recipe)
	})
}
