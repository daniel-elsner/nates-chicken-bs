package binding

import (
	"ncbs/api/model"
	"net/http"

	"github.com/labstack/echo/v4"
)

// RecipeBinderMiddleware is a middleware function that handles binding a recipe
// from the request body to a Recipe struct. If the binding fails, an error is
// returned.
func RecipeBinderMiddleware(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		recipe := new(model.Recipe)
		if err := c.Bind(recipe); err != nil {
			return c.JSON(http.StatusBadRequest, map[string]string{
				"error": "Failed to bind request body to Recipe struct",
			})
		}

		// could validate the recipe here

		c.Set("recipe", recipe)
		return next(c)
	}
}
