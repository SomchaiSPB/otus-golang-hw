package config

import (
	"fmt"
	"os"

	"gopkg.in/yaml.v3"
)

type Config struct {
	Logger LoggerConf
	App    AppConfig
}

type LoggerConf struct {
	File  string
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

	defer func() {
		cerr := f.Close()

		if cerr != nil {
			fmt.Println(cerr)
			return
		}
	}()

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
