package integration_tests

import (
	"app/pkg/server"
	"gorm.io/gorm"
	"net/http/httptest"
	"testing"
)

func setupTestCase(t *testing.T) (*httptest.Server, *gorm.DB, func()) {
	tx := testDB.Begin() // testDBはどこかで定義されているデータベースインスタンスです
	if tx.Error != nil {
		t.Fatalf("トランザクションの開始に失敗しました: %v", tx.Error)
	}

	testServer := httptest.NewServer(server.SetupServerWithDB(tx))

	return testServer, tx, func() {
		testServer.Close()
		tx.Rollback()
	}
}
