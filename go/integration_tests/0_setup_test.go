package integration_tests

import (
	"app/config"
	"app/models"
	"app/ogen"
	"app/pkg/db"
	"app/pkg/middleware"
	"app/pkg/repository"
	"app/services"
	"log"
	"net/http/httptest"
	"os"
	"testing"
)

var (
	testHTTP *httptest.Server
)

func TestMain(m *testing.M) {
	os.Setenv("ENV", "test")
	config.InitConfig()

	if err := db.InitDatabase(); err != nil {
		log.Fatalf("データベース接続エラー: %v", err)
	}
	log.Println("データベースに接続しました")

	srv := services.NewGoPracticeService(
		repository.NewGormMicropostRepository(db.DB),
		repository.NewGormUserRepository(db.DB),
	)

	httpServer, err := ogen.NewServer(srv)
	if err != nil {
		log.Fatalf("Failed to create server: %v", err)
	}

	wrappedServer := middleware.CorsMiddleware(httpServer)

	testHTTP = httptest.NewServer(wrappedServer)

	code := m.Run()

	testHTTP.Close()
	os.Exit(code)
}
