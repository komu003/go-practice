package config

import (
	"github.com/spf13/viper"
	"log"
	"os"
)

func InitConfig() {
	viper.AddConfigPath("config")
	viper.SetConfigType("yml")

	viper.SetConfigName("config")
	if err := viper.ReadInConfig(); err != nil {
		log.Fatalf("共通設定ファイルの読み込みに失敗しました: %v", err)
	}

	env := os.Getenv("ENV")
	if env == "" {
		env = "development"
	}
	viper.SetConfigName("config." + env)
	if err := viper.MergeInConfig(); err != nil {
		log.Fatalf("環境固有の設定ファイルの読み込みに失敗しました: %v", err)
	}
}
