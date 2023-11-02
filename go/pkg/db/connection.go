// パッケージ名をdbにして、他の部分から参照可能にする
package db

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"log"
	"os"
	"time"
)

// データベース接続情報の変数
var DB *gorm.DB

// InitDatabase 関数はデータベースの初期接続を行います。
func InitDatabase() error {
	dsn := "root:rootpassword@tcp(mysql)/mydatabase?charset=utf8mb4&parseTime=True&loc=Local"

	newLogger := logger.New(
		log.New(os.Stdout, "\r\n", log.LstdFlags),
		logger.Config{
			SlowThreshold: time.Second,
			LogLevel:      logger.Info,
			Colorful:      true,
		},
	)

	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{
		Logger: newLogger,
	})

	if err != nil {
		return err
	}

	DB = db

	return nil
}
