package integration_tests

import (
	"app/config"
	"app/pkg/server"
	"gorm.io/gorm"
	"os"
	"testing"
)

var testDB *gorm.DB

func TestMain(m *testing.M) {
	os.Setenv("ENV", "test")
	config.InitConfig()
	testDB = server.InitializeDatabase()

	code := m.Run()

	os.Exit(code)
}
