package route

import (
	"ncbs/api/handler"
	"ncbs/api/middleware/binding"
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

	e.GET("/recipes", func(c echo.Context) error {
		recipes, err := recipeHandler.GetAllRecipes()
		if err != nil {
			return handleError(err, c, "Failed to get recipes")
		}
		return c.JSON(http.StatusOK, recipes)
	})

	e.POST("/recipes", binding.RecipeBinderMiddleware(func(c echo.Context) error {
		recipe := c.Get("recipe").(*model.Recipe)
		err := recipeHandler.CreateRecipe(*recipe)
		if err != nil {
			return handleError(err, c, "Failed to create recipe")
		}

		return c.JSON(http.StatusOK, recipe)
	}))
}

func handleError(err error, c echo.Context, message string) error {
	c.Logger().Error(err)
	return echo.NewHTTPError(http.StatusInternalServerError, message)
}
