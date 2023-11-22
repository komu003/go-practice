package server

import (
	"app/config"
	"app/ogen"
	"app/pkg/db"
	"app/pkg/middleware"
	"app/pkg/repository"
	"app/services"
	"log"
	"net/http"
)

func SetupServer() http.Handler {
	InitializeConfigAndDatabase()

	srv := services.NewGoPracticeService(
		repository.NewGormMicropostRepository(db.DB),
		repository.NewGormUserRepository(db.DB),
	)

	httpServer, err := ogen.NewServer(srv)
	if err != nil {
		log.Fatalf("サーバーの作成に失敗しました: %v", err)
	}

	return middleware.CorsMiddleware(httpServer)
}

func InitializeConfigAndDatabase() {
	config.InitConfig()

	if err := db.InitDatabase(); err != nil {
		log.Fatalf("データベース接続エラー: %v", err)
	}
	log.Println("データベースに接続しました")
}
