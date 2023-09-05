package config

import (
	"log"

	"github.com/spf13/viper"
)

var EnvVar struct {
	ServerPort string `mapstructure:"SERVER_PORT"`
	DBUser     string `mapstructure:"DB_USER"`
	DBPassword string `mapstructure:"DB_PASSWORD"`
	DBHost     string `mapstructure:"DB_HOST"`
	DBPort     int    `mapstructure:"DB_PORT"`
	DBName     string `mapstructure:"DB_NAME"`
}

func setDefault() {
	viper.SetDefault("SERVER_PORT", "80")
}

func init() {
	viper.SetConfigFile(".env")
	viper.AutomaticEnv()

	setDefault()

	err := viper.ReadInConfig()
	if err != nil {
		log.Fatal("Failed to load .env file", err)
	}

	err = viper.Unmarshal(&EnvVar)
	if err != nil {
		log.Fatal("failed to unmarshal EnvVar", err)
	}
}
