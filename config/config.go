package config

import (
	"log"
	"os"

	"github.com/mcuadros/go-defaults"
	"github.com/spf13/viper"
)

type Config struct {
	DATABASE_URL string `mapstructure:"DATABASE_URL"`
	PORT         string `mapstructure:"PORT" default:"8080"`
	HOST         string `mapstructure:"HOST" default:"0.0.0.0"`
}

func Load() (*Config, error) {
	conf := &Config{}
	viper.SetConfigFile(".env")

	viper.AutomaticEnv()
	viper.BindEnv("DATABASE_URL", "DATABASE_URL")

	// Find and read the config file
	if err := viper.ReadInConfig(); err != nil {
		if os.IsNotExist(err) {
			log.Println("config file not found, reading vars from env")
		} else {
			return nil, err
		}
	}

	defaults.SetDefaults(conf)
	if err := viper.Unmarshal(conf); err != nil {
		return nil, err
	}
	return conf, nil
}
