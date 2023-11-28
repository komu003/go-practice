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
)

func TestAPIMicropostsGetIntegration(t *testing.T) {
	t.Skip("未実装")
}

func TestAPIMicropostsCountGetIntegration(t *testing.T) {
	cases := []struct {
		name     string
		setup    func(*gorm.DB) error
		expected int64
	}{
		{
			name: "ゼロマイクロポスト",
			setup: func(tx *gorm.DB) error {
				return nil
			},
			expected: 0,
		},
		{
			name: "ユーザーはいるがマイクロポストはゼロ",
			setup: func(tx *gorm.DB) error {
				user := models.User{Name: "ユーザー1", Email: "user1@example.com"}
				return tx.Create(&user).Error
			},
			expected: 0,
		},
		{
			name: "1ユーザー1マイクロポスト",
			setup: func(tx *gorm.DB) error {
				user := models.User{Name: "ユーザー2", Email: "user2@example.com"}
				if err := tx.Create(&user).Error; err != nil {
					return err
				}
				micropost := models.Micropost{Content: "マイクロポスト1", UserID: user.ID}
				return tx.Create(&micropost).Error
			},
			expected: 1,
		},
		{
			name: "1ユーザー複数マイクロポスト",
			setup: func(tx *gorm.DB) error {
				user := models.User{Name: "ユーザー3", Email: "user3@example.com"}
				if err := tx.Create(&user).Error; err != nil {
					return err
				}
				for i := 0; i < 5; i++ {
					micropost := models.Micropost{Content: fmt.Sprintf("マイクロポスト%d", i+1), UserID: user.ID}
					if err := tx.Create(&micropost).Error; err != nil {
						return err
					}
				}
				return nil
			},
			expected: 5,
		},
		{
			name: "複数ユーザー複数マイクロポスト",
			setup: func(tx *gorm.DB) error {
				for i := 0; i < 10; i++ {
					user := models.User{Name: fmt.Sprintf("ユーザー%d", i+4), Email: fmt.Sprintf("user%d@example.com", i+4)}
					if err := tx.Create(&user).Error; err != nil {
						return err
					}
					for j := 0; j < 3; j++ {
						micropost := models.Micropost{Content: fmt.Sprintf("ユーザー%dのマイクロポスト%d", i+4, j+1), UserID: user.ID}
						if err := tx.Create(&micropost).Error; err != nil {
							return err
						}
					}
				}
				return nil
			},
			expected: 30,
		},
	}

	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			testHTTP, tx, tearDown := setupTestCase(t)
			defer tearDown()

			require.NoError(t, tc.setup(tx), "テストデータのセットアップに失敗しました")

			url := fmt.Sprintf("%s/api/microposts/count", testHTTP.URL)
			res, err := http.Get(url)
			require.NoError(t, err, "HTTPリクエストの送信に失敗しました")
			defer res.Body.Close()

			var countResponse struct {
				Count int64 `json:"count"`
			}
			err = json.NewDecoder(res.Body).Decode(&countResponse)
			require.NoError(t, err, "レスポンスのデコードに失敗しました")

			assert.Equal(t, tc.expected, countResponse.Count, "マイクロポスト数が期待と異なります")
		})
	}
}
