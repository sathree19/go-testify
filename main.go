package main

import (
	"go-testify/cafe"
	"net/http"
)

func main() {
	http.HandleFunc(`/`, cafe.MainHandle)
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		panic(err)
	}
}
