package config

import (
	"github.com/spf13/viper"
)

type Config struct {
	ApiHost string
	ApiPort int
}

func NewConfig() *Config {
	viper.AutomaticEnv()
	return &Config{
		ApiHost: viper.GetString("HOST"),
		ApiPort: viper.GetInt("PORT"),
	}
}
