package main

import (
	"fmt"
	"net/http"
	"data"
	"handlers"
)

func main() {
	store := data.NewUrlStore()
	http.HandleFunc("/shorten", handlers.ShortenURLHandler(store))
	http.HandleFunc("/", handlers.RedirectHandler(store))

	fmt.Println("Server started on :8080")
	http.ListenAndServe(":8080", nil)
}
