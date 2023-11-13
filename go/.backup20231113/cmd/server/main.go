package main

import (
	"app/handlers"
	"app/pkg/db"
	"log"
	"net/http"
)

func main() {
	if err := db.InitDatabase(); err != nil {
		log.Fatalf("データベース接続エラー: %v", err)
	}

	http.HandleFunc("/api/users/count", handlers.UsersCountHandler)
	http.HandleFunc("/api/microposts/count", handlers.MicropostsCountHandler)

	log.Fatal(http.ListenAndServe(":8080", nil))
}
