package server

import (
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
	dbInstance := InitializeDatabase()
	return SetupServerWithDB(dbInstance)
}

func SetupServerWithDB(dbInstance *gorm.DB) http.Handler {
	micropostRepository := repository.NewGormMicropostRepository(dbInstance)
	userRepository := repository.NewGormUserRepository(dbInstance)

	goPracticeService := services.NewGoPracticeService(micropostRepository, userRepository)

	httpServer, err := ogen.NewServer(goPracticeService)
	if err != nil {
		log.Fatalf("サーバーの作成に失敗しました: %v", err)
	}

	return middleware.CorsMiddleware(httpServer)
}

func InitializeDatabase() *gorm.DB {
	dbInstance, err := db.InitDatabase()
	if err != nil {
		log.Fatalf("データベース接続エラー: %v", err)
	}
	log.Println("データベースに接続しました")

	return dbInstance
}
