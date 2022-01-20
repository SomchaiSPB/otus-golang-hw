package config

import (
	"fmt"
	"os"

	"gopkg.in/yaml.v3"
)

type Config struct {
	Logger LoggerConf
	App    AppConfig
	DB     DBConfig
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

type DBConfig struct {
	Dsn string
}

func NewConfig(configPath string) (*Config, error) {
	var cfg *Config

	f, err := os.Open(configPath)
	if err != nil {
		fmt.Println(err)
		return cfg, err
	}

	defer func() {
		cerr := f.Close()

		if cerr != nil {
			fmt.Println(cerr)
			return
		}
	}()

	decoder := yaml.NewDecoder(f)
	err = decoder.Decode(&cfg)

	if err != nil {
		fmt.Println(err)
	}

	return cfg, nil
}
