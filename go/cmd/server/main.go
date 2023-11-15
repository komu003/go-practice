package main

import (
    "context"
    "log"
    "net/http"
	"errors"
    "app/ogen"
    "app/pkg/db"
    "app/models"
)

type goPracticeService struct {}

func (s *goPracticeService) APIMicropostsCountGet(ctx context.Context) (ogen.APIMicropostsCountGetRes, error) {
	if db.DB == nil {
        return nil, errors.New("データベース接続が利用できません")
    }
	var count int64
    if err := db.DB.Model(&models.Micropost{}).Count(&count).Error; err != nil {
        return &ogen.APIMicropostsCountGetInternalServerError{}, err
    }
    return &ogen.CountResponse{Count: ogen.NewOptInt(int(count))}, nil
}

func (s *goPracticeService) APIUsersCountGet(ctx context.Context) (ogen.APIUsersCountGetRes, error) {
	if db.DB == nil {
        return nil, errors.New("データベース接続が利用できません")
    }
	var count int64
    if err := db.DB.Model(&models.User{}).Count(&count).Error; err != nil {
        return &ogen.APIUsersCountGetInternalServerError{}, err
    }
    return &ogen.CountResponse{Count: ogen.NewOptInt(int(count))}, nil
}

func main() {
	if err := db.InitDatabase(); err != nil {
		log.Fatalf("データベース接続エラー: %v", err)
	}

    srv := &goPracticeService{}
    httpServer, err := ogen.NewServer(srv)
    if err != nil {
        log.Fatalf("Failed to create server: %v", err)
    }

    log.Println("Server is running at http://localhost:8080")
    log.Fatal(http.ListenAndServe(":8080", httpServer))
}
