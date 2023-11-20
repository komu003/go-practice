// setup_test.go
package integration_tests

import (
	"app/models"
	"log"
	"os"
	"testing"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var db *gorm.DB

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

	code := m.Run()

	os.Exit(code)
}
