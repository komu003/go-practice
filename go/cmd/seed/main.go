package main

import (
	"app/config"
	"app/models"
	"app/pkg/db"
	"fmt"
	"log"
)

func main() {
	config.InitConfig()

	if err := db.InitDatabase(); err != nil {
		log.Fatalf("データベース接続エラー: %v", err)
	}
	fmt.Println("データベースに接続しました")

	var users []models.User
	for i := 1; i <= 3; i++ {
		user := models.User{
			Name:  fmt.Sprintf("User%d", i),
			Email: fmt.Sprintf("user%d@example.com", i),
		}
		users = append(users, user)
	}

	db.DB.Create(&users)

	var microposts []models.Micropost
	for _, user := range users {
		for j := 1; j <= 3; j++ {
			micropost := models.Micropost{
				Content: fmt.Sprintf("Micropost %d by User %d", j, user.ID),
				UserID:  user.ID,
			}
			microposts = append(microposts, micropost)
		}
	}

	db.DB.Create(&microposts)

	fmt.Println("シードデータの作成が完了しました。")
}
