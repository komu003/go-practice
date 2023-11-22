package integration_tests

import (
	"app/models"
	"app/pkg/db"
	"encoding/json"
	"fmt"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"net/http"
	"testing"
)

func TestAPIUsersGetIntegration(t *testing.T) {
	t.Skip("未実装")
}

func TestAPIUsersCountGetIntegration(t *testing.T) {
	cases := []struct {
		name     string
		setup    func() error
		expected int64
	}{
		{
			name: "ゼロユーザー",
			setup: func() error {
				return nil
			},
			expected: 0,
		},
		{
			name: "1ユーザー",
			setup: func() error {
				return db.DB.Create(&models.User{Name: "テストユーザー1", Email: "test1@example.com"}).Error
			},
			expected: 1,
		},
	}

	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			require.NoError(t, tc.setup(), "テストデータのセットアップに失敗しました")

			url := fmt.Sprintf("%s/api/users/count", testHTTP.URL)
			res, err := http.Get(url)
			require.NoError(t, err, "HTTPリクエストの送信に失敗しました")
			defer res.Body.Close()

			var countResponse struct {
				Count int64 `json:"count"`
			}

			err = json.NewDecoder(res.Body).Decode(&countResponse)
			require.NoError(t, err, "レスポンスのデコードに失敗しました")

			assert.Equal(t, http.StatusOK, res.StatusCode, "HTTPステータスコードが期待と異なります")
			assert.Equal(t, tc.expected, countResponse.Count, "ユーザー数が期待と異なります")
		})
	}
}
