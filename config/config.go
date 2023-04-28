package config

import (
	"log"

	"github.com/spf13/viper"
)

type Config struct {
	DATABASE_URL string `mapstructure:"DATABASE_URL"`
	PORT         string `mapstructure:"PORT"`
	HOST         string `mapstructure:"HOST"`
}

func Load() *Config {
	conf := &Config{}
	viper.SetConfigFile(".env")
	// Find and read the config file
	if err := viper.ReadInConfig(); err != nil {
		log.Fatalf("Error while reading config file %s", err)
	}
	if err := viper.Unmarshal(conf); err != nil {
		log.Fatalf("Error while loading config to struct: %s", err)
	}
	return conf
}
