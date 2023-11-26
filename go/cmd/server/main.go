package main

import (
	"app/config"
	"app/pkg/server"
	"log"
	"net/http"
)

func main() {
	config.InitConfig()
	server := server.SetupServer()
	log.Println("サーバー実行中：http://localhost:8080")
	log.Fatal(http.ListenAndServe(":8080", server))
}
