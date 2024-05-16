package main

import (
	"fmt"
	"moneylion-conding-challenge/api/handlers"
	"net/http"
)

func main() {
	mux := http.NewServeMux()

	mux.HandleFunc("GET /feed", handlers.FeedHandler)

	if err := http.ListenAndServe(":8080", mux); err != nil {
		fmt.Println(err.Error())
	}
}
