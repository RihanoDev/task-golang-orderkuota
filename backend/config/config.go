package config

import (
	"log"
	"sync"

	"github.com/joho/godotenv"
	"github.com/spf13/viper"
)

var one sync.Once

func LoadENV() {
	one.Do(func() {
		err := godotenv.Load(".env")
		if err != nil {
			log.Println("Warning: .env file not found, using system environment variable")
		}
		viper.AutomaticEnv()
	})
}

func GetConfig(key string) string {
	LoadENV()
	return viper.GetString(key)
}
