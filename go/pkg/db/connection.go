// パッケージ名をdbにして、他の部分から参照可能にする
package db

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

// データベース接続情報の変数
var DB *gorm.DB

// InitDatabase 関数はデータベースの初期接続を行います。
func InitDatabase() error {
	dsn := "root:rootpassword@tcp(mysql)/mydatabase?charset=utf8mb4&parseTime=True&loc=Local"

	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		return err
	}
	DB = db

	return nil
}
