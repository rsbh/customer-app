package config

import (
	"github.com/spf13/viper"
)

type Config struct {
	DATABASE_URL string `mapstructure:"DATABASE_URL"`
	PORT         string `mapstructure:"PORT"`
	HOST         string `mapstructure:"HOST"`
}

func Load() (*Config, error) {
	conf := &Config{}
	viper.SetConfigFile(".env")
	// Find and read the config file
	if err := viper.ReadInConfig(); err != nil {
		return nil, err
	}
	if err := viper.Unmarshal(conf); err != nil {
		return nil, err
	}
	return conf, nil
}
