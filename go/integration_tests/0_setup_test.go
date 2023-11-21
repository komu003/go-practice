package integration_tests

import (
	"app/models"
	"app/ogen"
	"app/pkg/middleware"
	"app/pkg/repository"
	"app/services"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"log"
	"net/http/httptest"
	"os"
	"testing"
)

var (
	db       *gorm.DB
	testHTTP *httptest.Server
)

func TestMain(m *testing.M) {
	var err error

	dsn := "root:rootpassword@tcp(mysql)/test_mydatabase?charset=utf8mb4&parseTime=True&loc=Local"
	db, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalf("データベースへの接続に失敗しました: %v", err)
	}

	err = db.AutoMigrate(&models.User{})
	if err != nil {
		log.Fatalf("マイグレーションに失敗しました: %v", err)
	}

	srv := services.NewGoPracticeService(
		repository.NewGormMicropostRepository(db),
		repository.NewGormUserRepository(db),
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
