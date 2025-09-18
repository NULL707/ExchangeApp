package config

import (
	"log"

	"github.com/spf13/viper"
)

type Config struct {
	App struct {
		Name string `yaml:"name"`
		Port string `yaml:"port"`
	} `yaml:"app"`
	Database struct {
		// Host     string `yaml:"host"`
		// Port     string `yaml:"port"`
		// User     string `yaml:"user"`
		// Password string `yaml:"password"`
		// Name     string `yaml:"name"`
		Dsn          string `yaml:"dsn"`
		MaxIdleConns int    `yaml:"MaxIdleConns"`
		MaxOpenConns int    `yaml:"MaxOpenConns"`
	} `yaml:"database"`
}

var AppConfig *Config

func InitConfig() {
	viper.SetConfigName("config")
	viper.SetConfigType("yaml")
	viper.AddConfigPath("./config")
	if err := viper.ReadInConfig(); err != nil {
		log.Fatalf("Error reading config file, %s", err)
	}
	AppConfig = &Config{}
	if err := viper.Unmarshal(AppConfig); err != nil {
		log.Fatalf("Error unmarshal config file, %s", err)
	}
	InitDB()
	InitRedis()
}
