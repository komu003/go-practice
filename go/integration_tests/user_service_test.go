package integration_tests

import (
	"app/models"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestAPIUsersGetIntegration(t *testing.T) {
	t.Skip("未実装")
}

func TestAPIUsersCountGetIntegration(t *testing.T) {
	err := db.Create(&models.User{Name: "テストユーザー1", Email: "test1@example.com"}).Error
	if err != nil {
		t.Fatalf("テストデータの挿入に失敗しました: %v", err)
	}

	var count int64
	err = db.Model(&models.User{}).Count(&count).Error
	if err != nil {
		t.Fatalf("データのカウントに失敗しました: %v", err)
	}
	assert.Equal(t, int64(1), count, "ユーザー数が期待と異なります")

	t.Skip("実装不十分")
}
