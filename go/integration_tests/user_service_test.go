package integration_tests

import (
	"app/models"
	"encoding/json"
	"fmt"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"gorm.io/gorm"
	"net/http"
	"testing"
	"time"
)

func TestAPIUsersGetIntegration(t *testing.T) {
	cases := []struct {
		name     string
		setup    func(*gorm.DB) error
		expected []models.User
	}{
		{
			name: "ゼロユーザー",
			setup: func(tx *gorm.DB) error {
				// 何もしない（ユーザーを作成しない）
				return nil
			},
			expected: []models.User{},
		},
		{
			name: "1ユーザー",
			setup: func(tx *gorm.DB) error {
				// 1ユーザーを作成
				user := models.User{Name: "テストユーザー1", Email: "test1@example.com"}
				return tx.Create(&user).Error
			},
			expected: []models.User{
				{Name: "テストユーザー1", Email: "test1@example.com"},
			},
		},
	}

	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			testHTTP, tx, tearDown := setupTestCase(t)
			defer tearDown()

			require.NoError(t, tc.setup(tx), "テストデータのセットアップに失敗しました")

			url := fmt.Sprintf("%s/api/users", testHTTP.URL)
			res, err := http.Get(url)
			require.NoError(t, err, "HTTPリクエストの送信に失敗しました")
			defer res.Body.Close()

			var usersResponse []struct {
				ID        int       `json:"id"`
				Name      string    `json:"name"`
				Email     string    `json:"email"`
				CreatedAt time.Time `json:"createdAt"`
				UpdatedAt time.Time `json:"updatedAt"`
			}
			err = json.NewDecoder(res.Body).Decode(&usersResponse)
			require.NoError(t, err, "レスポンスのデコードに失敗しました")

			assert.Equal(t, len(tc.expected), len(usersResponse), "レスポンスのユーザー数が期待と異なります")

			for i, user := range usersResponse {
				assert.Equal(t, tc.expected[i].Name, user.Name, "ユーザー名が期待と異なります")
				assert.Equal(t, tc.expected[i].Email, user.Email, "ユーザーメールが期待と異なります")
			}
		})
	}
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
				users := make([]models.User, 100)
				for i := 0; i < 100; i++ {
					users[i] = models.User{Name: fmt.Sprintf("テストユーザー%d", i+1), Email: fmt.Sprintf("test%d@example.com", i+1)}
				}
				return tx.Create(&users).Error
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
