package config

import (
	"github.com/joho/godotenv"
	"github.com/spf13/viper"
)

type Config struct {
	POSTGRES_USERNAME string `mapstructure:"POSTGRES_USERNAME"`
	POSTGRES_PASSWORD string `mapstructure:"POSTGRES_PASSWORD"`
	POSTGRES_HOST     string `mapstructure:"POSTGRES_HOST"`
	POSTGRES_PORT     string `mapstructure:"POSTGRES_PORT"`
	DB_NAME           string `mapstructure:"DB_NAME"`
	SERVER_ADDRESS    string `mapstructure:"SERVER_ADDRESS"`
}

func LoadConfig() (config Config, err error) {
	err = godotenv.Load()
	if err != nil {
		return
	}
	viper.SetConfigType("env")

	viper.AutomaticEnv()

	config.DB_NAME = viper.GetString("DB_NAME")
	config.POSTGRES_USERNAME = viper.GetString("POSTGRES_USERNAME")
	config.POSTGRES_PASSWORD = viper.GetString("POSTGRES_PASSWORD")
	config.POSTGRES_HOST = viper.GetString("POSTGRES_HOST")
	config.POSTGRES_PORT = viper.GetString("POSTGRES_PORT")
	config.SERVER_ADDRESS = viper.GetString("SERVER_ADDRESS")
	return
}
