package integration_tests

import (
	"app/models"
	"app/ogen"
	"app/pkg/repository"
	"app/services"
	"context"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"testing"
)

func TestAPIUsersGetIntegration(t *testing.T) {
	t.Skip("未実装")
}

func TestAPIUsersCountGetIntegration(t *testing.T) {
	userRepo := repository.NewGormUserRepository(db)
	userService := services.NewUserService(userRepo)

	testUser := &models.User{Name: "テストユーザー1", Email: "test1@example.com"}
	err := db.Create(testUser).Error
	require.NoError(t, err, "テストデータの挿入に失敗しました")

	defer db.Delete(testUser)

	ctx := context.Background()
	res, err := userService.APIUsersCountGet(ctx)
	require.NoError(t, err, "APIUsersCountGetの呼び出しに失敗しました")

	countResponse, ok := res.(*ogen.CountResponse)
	require.True(t, ok, "レスポンスタイプが期待と異なります")

	assert.Equal(t, int64(1), int64(countResponse.Count.Value), "ユーザー数が期待と異なります")
}
