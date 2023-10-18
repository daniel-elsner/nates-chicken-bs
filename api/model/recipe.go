package model

type Recipe struct {
	ID        string   `json:"id"`
	Name      string   `json:"name"`
	PrepSteps []string `json:"prepSteps"`
	CookSteps []string `json:"cookSteps"`
}
