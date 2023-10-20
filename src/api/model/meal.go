package model

type Meal struct {
	ID        string   `json:"id"`
	Name      string   `json:"name"`
	RecipeIDs []string `json:"recipeIds"`
	PrepSteps []Step   `json:"prepSteps"`
	CookSteps []Step   `json:"cookSteps"`
}
