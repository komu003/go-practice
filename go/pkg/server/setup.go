package server

import (
	"app/config"
	"app/ogen"
	"app/pkg/db"
	"app/pkg/middleware"
	"app/pkg/repository"
	"app/services"
	"gorm.io/gorm"
	"log"
	"net/http"
)

func SetupServer() http.Handler {
	dbInstance := InitializeConfigAndDatabase()

	srv := services.NewGoPracticeService(
		repository.NewGormMicropostRepository(dbInstance),
		repository.NewGormUserRepository(dbInstance),
	)

	httpServer, err := ogen.NewServer(srv)
	if err != nil {
		log.Fatalf("サーバーの作成に失敗しました: %v", err)
	}

	return middleware.CorsMiddleware(httpServer)
}

func InitializeConfigAndDatabase() *gorm.DB {
	config.InitConfig()

	dbInstance, err := db.InitDatabase()
	if err != nil {
		log.Fatalf("データベース接続エラー: %v", err)
	}
	log.Println("データベースに接続しました")

	return dbInstance
}
