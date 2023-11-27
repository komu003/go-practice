package integration_tests

import (
	"app/models"
	"app/pkg/server"
	"encoding/json"
	"fmt"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"gorm.io/gorm"
	"net/http"
	"net/http/httptest"
	"testing"
)

func setupTestCase(t *testing.T) (*httptest.Server, *gorm.DB, func()) {
	tx := testDB.Begin()
	if tx.Error != nil {
		t.Fatalf("トランザクションの開始に失敗しました: %v", tx.Error)
	}

	testServer := httptest.NewServer(server.SetupServerWithDB(tx))

	return testServer, tx, func() {
		testServer.Close()
		tx.Rollback()
	}
}

func TestAPIUsersGetIntegration(t *testing.T) {
	t.Skip("未実装")
}

func TestAPIUsersCountGetIntegration(t *testing.T) {
	cases := []struct {
		name     string
		setup    func(*gorm.DB) error
		expected int64
	}{
		{
			name: "ゼロユーザー",
			setup: func(tx *gorm.DB) error {
				return nil
			},
			expected: 0,
		},
		{
			name: "1ユーザー",
			setup: func(tx *gorm.DB) error {
				return tx.Create(&models.User{Name: "テストユーザー1", Email: "test1@example.com"}).Error
			},
			expected: 1,
		},
		{
			name: "100ユーザー",
			setup: func(tx *gorm.DB) error {
				for i := 0; i < 100; i++ {
					user := models.User{Name: fmt.Sprintf("テストユーザー%d", i+1), Email: fmt.Sprintf("test%d@example.com", i+1)}
					if err := tx.Create(&user).Error; err != nil {
						return err
					}
				}
				return nil
			},
			expected: 100,
		},
	}

	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			testHTTP, tx, tearDown := setupTestCase(t)
			defer tearDown()

			require.NoError(t, tc.setup(tx), "テストデータのセットアップに失敗しました")

			url := fmt.Sprintf("%s/api/users/count", testHTTP.URL)
			res, err := http.Get(url)
			require.NoError(t, err, "HTTPリクエストの送信に失敗しました")
			defer res.Body.Close()

			var countResponse struct {
				Count int64 `json:"count"`
			}
			err = json.NewDecoder(res.Body).Decode(&countResponse)
			require.NoError(t, err, "レスポンスのデコードに失敗しました")

			assert.Equal(t, tc.expected, countResponse.Count, "ユーザー数が期待と異なります")
		})
	}
}
