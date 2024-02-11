package controllers

import (
	"fmt"
	"net/http"
)

func StartMainServer() error {
	http.HandleFunc("/", top)
	fmt.Println("Server is running on port 8080...")
	return http.ListenAndServe(":8080", nil)
}
