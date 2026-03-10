package config

import (
	"os"

	"github.com/ilyakaznacheev/cleanenv"
)

type Config struct {
	Environment string `yaml:"environment"`
	JWT      JWTConfig      `yaml:"jwt"`
	Database DatabaseConfig `yaml:"database"`
	GRPC     GRPCConfig     `yaml:"grpc"`
	DataBaseUrl string `yaml:"database_url"`
}

type JWTConfig struct {
	Secret	 string `yaml:"secret"`
	Expiration int    `yaml:"expiration"`
}

type DatabaseConfig struct {
	Host     string `yaml:"host"`
	Port     int    `yaml:"port"`
	User     string `yaml:"user"`
	Password string `yaml:"password"`
	DBName   string `yaml:"dbname"`
}

type GRPCConfig struct {
	Port int `yaml:"port"`
}


func MustLoadConfig() *Config {
	ConfigPath  := os.Getenv("CONFIG_AUTH_PATH") 
	if ConfigPath == "" {
		panic("CONFIG_AUTH_PATH is not set")
	}

	var cfg Config

	err := cleanenv.ReadConfig(ConfigPath, &cfg)
	if err != nil {
		panic(err)
	}

	return &cfg
}