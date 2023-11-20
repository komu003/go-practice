// setup_test.go
package integration_tests

import (
	"app/models"
	"app/ogen"
	"app/pkg/middleware"
	"app/pkg/repository"
	"app/services"
	"gorm.io/driver/sqlite"
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

type GoPracticeService struct {
	*services.MicropostService
	*services.UserService
}

func TestMain(m *testing.M) {
	var err error
	db, err = gorm.Open(sqlite.Open(":memory:"), &gorm.Config{})
	if err != nil {
		log.Fatalf("データベースへの接続に失敗しました: %v", err)
	}

	err = db.AutoMigrate(&models.User{})
	if err != nil {
		log.Fatalf("マイグレーションに失敗しました: %v", err)
	}

	srv := &GoPracticeService{
		MicropostService: services.NewMicropostService(repository.NewGormMicropostRepository(db)),
		UserService:      services.NewUserService(repository.NewGormUserRepository(db)),
	}

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
