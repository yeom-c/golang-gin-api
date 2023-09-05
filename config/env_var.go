package config

import (
	"log"

	"github.com/spf13/viper"
)

var EnvVar struct {
	ServerPort string `mapstructure:"SERVER_PORT" default:"8080"`
}

func init() {
	viper.SetConfigFile(".env")
	viper.AutomaticEnv()

	err := viper.ReadInConfig()
	if err != nil {
		log.Fatal("Failed to load .env file", err)
	}

	err = viper.Unmarshal(&EnvVar)
	if err != nil {
		log.Fatal("failed to unmarshal EnvVar", err)
	}
}
