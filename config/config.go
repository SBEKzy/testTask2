package config

import (
	"github.com/spf13/viper"
	"log"
)

type Config struct {
	DBHost string `mapstructure:"DB_HOST"`
	Port   string `mapstructure:"PORT"`
	DBName string `mapstructure:"DB_NAME"`
}

func New() *Config {
	var config Config
	viper.AddConfigPath("./")
	viper.SetConfigName("config")
	viper.SetConfigType("env")
	err := viper.ReadInConfig()
	if err != nil {
		log.Fatalln(err)
	}
	err = viper.Unmarshal(&config)
	if err != nil {
		log.Fatalln(err)
	}

	return &config
}
