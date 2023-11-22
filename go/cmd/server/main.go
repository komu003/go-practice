package main

import (
	"app/pkg/server"
	"log"
	"net/http"
)

func main() {
	server := server.SetupServer()
	log.Println("サーバー実行中：http://localhost:8080")
	log.Fatal(http.ListenAndServe(":8080", server))
}
