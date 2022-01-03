package config

import (
	"fmt"
	"gopkg.in/yaml.v3"
	"os"
)

type Config struct {
	Logger LoggerConf
	App    AppConfig
}

type LoggerConf struct {
	Level string
}

type AppConfig struct {
	Port    string
	Storage string
}

func NewConfig(configPath string) *Config {
	var cfg Config

	f, err := os.Open(configPath)
	if err != nil {
		fmt.Println(err)
	}
	defer f.Close()

	decoder := yaml.NewDecoder(f)
	err = decoder.Decode(&cfg)

	if err != nil {
		fmt.Println(err)
	}

	return &cfg
}
