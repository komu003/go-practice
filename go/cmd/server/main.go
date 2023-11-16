package main

import (
	"app/pkg/db"
	"app/services"
	"app/ogen"
	"log"
	"net/http"
)

// Handler aggregates all service handlers
type GoPracticeService struct {
	*services.MicropostService
	*services.UserService
}

func main() {
	if err := db.InitDatabase(); err != nil {
		log.Fatalf("データベース接続エラー: %v", err)
	}

	srv := &GoPracticeService{
		MicropostService: &services.MicropostService{},
		UserService:      &services.UserService{},
	}

	httpServer, err := ogen.NewServer(srv)
	log.Printf("%T\n",httpServer)
	if err != nil {
		log.Fatalf("Failed to create server: %v", err)
	}

	log.Println("Server is running at http://localhost:8080")
	log.Fatal(http.ListenAndServe(":8080", httpServer))
}
