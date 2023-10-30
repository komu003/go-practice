package main

import (
	"app/handlers"
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/api/hello", handlers.HelloHandler)

	log.Fatal(http.ListenAndServe(":8080", nil))
}
