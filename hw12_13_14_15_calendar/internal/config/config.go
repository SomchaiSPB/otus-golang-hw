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
	Host    string
	Port    string
	Storage string
}

func NewConfig(configPath string) (*Config, error) {
	var cfg *Config

	f, err := os.Open(configPath)
	defer f.Close()

	if err != nil {
		fmt.Println(err)
		return cfg, err
	}

	decoder := yaml.NewDecoder(f)
	err = decoder.Decode(&cfg)

	if err != nil {
		fmt.Println(err)
	}

	return cfg, nil
}
