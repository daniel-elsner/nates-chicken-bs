package model

type Step struct {
	Description string   `json:"description"`
	Duration    int      `json:"duration"` // Duration in seconds
	Ingredients []string `json:"ingredients"`
}
