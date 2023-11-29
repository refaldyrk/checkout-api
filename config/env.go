package config

import (
	"log"

	"github.com/joho/godotenv"
	"github.com/spf13/viper"
)

type EnvsSchema struct {
	ADDR       string
	JWT_SECRET string
	JWT_EXP    int
	DB_URI     string
}

var Envs *EnvsSchema

func init() {
	err := godotenv.Load()
	if err != nil {
		log.Println("Error loading .env file")
	}

	viper.SetConfigType("env")
	viper.AutomaticEnv()

	Envs = &EnvsSchema{
		ADDR:       viper.GetString("ADDR"),
		JWT_SECRET: viper.GetString("JWT_SECRET"),
		JWT_EXP:    viper.GetInt("JWT_EXP"),
		DB_URI:     viper.GetString("DB_URI"),
	}
}
