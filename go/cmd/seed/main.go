package main

import (
	"app/config"
	"app/models"
	"app/pkg/server"
	"fmt"
)

func main() {
	config.InitConfig()
	db := server.InitializeDatabase()

	var users []models.User
	for i := 1; i <= 3; i++ {
		user := models.User{
			Name:  fmt.Sprintf("User%d", i),
			Email: fmt.Sprintf("user%d@example.com", i),
		}
		users = append(users, user)
	}

	db.Create(&users)

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

	db.Create(&microposts)

	fmt.Println("シードデータの作成が完了しました。")
}
