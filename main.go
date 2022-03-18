package main

import (
	"net/http"

	"github.com/sumitroajiprabowo/go-oauth2/controllers"
)

func main() {
	http.HandleFunc("/google/login", controllers.GoogleLogin)
	http.HandleFunc("/google/callback", controllers.GoogleCallback)

	http.ListenAndServe(":8000", nil)
}

// Language: go
