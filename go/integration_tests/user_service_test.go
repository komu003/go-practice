package integration_tests

import (
	"app/models"
	"encoding/json"
	"fmt"
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestAPIUsersGetIntegration(t *testing.T) {
	t.Skip("未実装")
}

func TestAPIUsersCountGetIntegration(t *testing.T) {
	testUser := &models.User{Name: "テストユーザー1", Email: "test1@example.com"}
	err := db.Create(testUser).Error
	require.NoError(t, err, "テストデータの挿入に失敗しました")
	defer db.Delete(testUser)

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
	assert.Equal(t, int64(1), countResponse.Count, "ユーザー数が期待と異なります")
}
