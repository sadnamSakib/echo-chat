package config

import (
	"log"

	"github.com/spf13/viper"
)

type Config struct {
	Database struct {
		MongoDB struct {
			URI      string
			Database string
		}
	}
	JWT struct {
		Secret string
	}
}

var AppConfig Config

func LoadConfig() {
	viper.SetConfigName("config")
	viper.SetConfigType("yaml")
	viper.AddConfigPath("configs")

	if err := viper.ReadInConfig(); err != nil {
		log.Fatalf("Error reading config file: %s", err)
	}

	if err := viper.Unmarshal(&AppConfig); err != nil {
		log.Fatalf("Unable to decode config into struct: %s", err)
	}
}
