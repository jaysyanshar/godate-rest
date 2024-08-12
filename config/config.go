package config

import (
	"log"

	"github.com/spf13/viper"
)

type Config struct {
	// application
	AppName string `mapstructure:"APP_NAME"`
	AppPort string `mapstructure:"APP_PORT"`

	// database
	DbDriver   string `mapstructure:"DB_DRIVER"`
	DbHost     string `mapstructure:"DB_HOST"`
	DbPort     string `mapstructure:"DB_PORT"`
	DbUser     string `mapstructure:"DB_USER"`
	DbPassword string `mapstructure:"DB_PASSWORD"`
	DbName     string `mapstructure:"DB_NAME"`

	// auth
	JwtSecret string `mapstructure:"JWT_SECRET"`
}

var config *Config

func init() {
	viper.SetConfigFile(".env")
	viper.AutomaticEnv()

	if err := viper.ReadInConfig(); err != nil {
		log.Println("Error reading .env file, using environment variables")
	}

	config = &Config{}
	if err := viper.Unmarshal(config); err != nil {
		log.Fatalf("Unable to decode into struct, %v", err)
	}
}

func Get() *Config {
	return config
}
