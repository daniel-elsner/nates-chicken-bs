package binding

import (
	"ncbs/api/model"
	"net/http"

	"github.com/labstack/echo/v4"
)

// RecipeBinderMiddleware is a middleware function that handles binding a recipe
// from the request body to a valid Recipe struct.
func RecipeBinderMiddleware(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		recipe := new(model.Recipe)
		if err := c.Bind(recipe); err != nil {
			return echo.NewHTTPError(http.StatusBadRequest, "Failed to bind request body to Recipe struct")
		}

		if recipe.Name == "" {
			return echo.NewHTTPError(http.StatusBadRequest, "Recipe name cannot be empty")
		}

		if recipe.PrepSteps == nil || len(recipe.PrepSteps) == 0 {
			return echo.NewHTTPError(http.StatusBadRequest, "Recipe prep steps cannot be empty")
		}

		if recipe.CookSteps == nil || len(recipe.CookSteps) == 0 {
			return echo.NewHTTPError(http.StatusBadRequest, "Recipe cook steps cannot be empty")
		}

		c.Set("recipe", recipe)
		return next(c)
	}
}
