// Package config Description: This file contains the configuration loader for the application using Viper.
package config

import "github.com/spf13/viper"

type Config struct {
	AppPort string

	DBUrl string

	JWTSecret string

	AdminAPIKey string // master key for managing API keys (from env)
}

func Load() (*Config, error) {
	viper.SetConfigFile(".env")
	viper.AutomaticEnv()

	if err := viper.ReadInConfig(); err != nil {
		return nil, err
	}

	return &Config{
		AppPort:     viper.GetString("APP_PORT"),
		DBUrl:       viper.GetString("DB_URL"),
		JWTSecret:   viper.GetString("JWT_SECRET"),
		AdminAPIKey: viper.GetString("ADMIN_API_KEY"),
	}, nil
}
