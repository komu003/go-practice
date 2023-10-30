package main

import (
	"log"
	"net/http"
	"app/handlers"
)

func main() {
	http.HandleFunc("/api/hello", handlers.HelloHandler)

	log.Fatal(http.ListenAndServe(":8080", nil))
}
