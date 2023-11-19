package main

import (
	"app/ogen"
	"app/pkg/db"
	"app/pkg/middleware"
	"app/pkg/repository"
	"app/services"
	"log"
	"net/http"
)

type GoPracticeService struct {
	*services.MicropostService
	*services.UserService
}

func main() {
	if err := db.InitDatabase(); err != nil {
		log.Fatalf("データベース接続エラー: %v", err)
	}

	srv := &GoPracticeService{
		MicropostService: services.NewMicropostService(repository.NewGormMicropostRepository()),
		UserService:      services.NewUserService(repository.NewGormUserRepository()),
	}

	httpServer, err := ogen.NewServer(srv)
	if err != nil {
		log.Fatalf("Failed to create server: %v", err)
	}

	wrappedServer := middleware.CorsMiddleware(httpServer)

	log.Println("Server is running at http://localhost:8080")
	log.Fatal(http.ListenAndServe(":8080", wrappedServer))
}
