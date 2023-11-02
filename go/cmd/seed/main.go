package main

import (
	"app/models"
	"app/pkg/db"
	"fmt"
	"log"
)

func main() {
	// データベースに接続
	if err := db.InitDatabase(); err != nil {
		log.Fatalf("データベース接続エラー: %v", err)
	}
	fmt.Println("データベースに接続しました")

	for i := 1; i <= 3; i++ {
		user := models.User{
			Name:  fmt.Sprintf("User%d", i),
			Email: fmt.Sprintf("user%d@example.com", i),
		}
		db.DB.Create(&user)

		for j := 1; j <= 3; j++ {
			micropost := models.Micropost{
				Content: fmt.Sprintf("Micropost %d by User %d", j, i),
				UserID:  user.ID,
			}
			db.DB.Create(&micropost)
		}
	}

	fmt.Println("シードデータの作成が完了しました。")
}
