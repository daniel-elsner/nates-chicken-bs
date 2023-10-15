package main

import (
	"fmt"
	"net/http"
)

func recipeHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Nates chicken bullshit")
}

func main() {
	http.HandleFunc("/recipe", recipeHandler)
	http.ListenAndServe(":8080", nil)
}
