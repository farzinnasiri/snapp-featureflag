package config

import (
	"gopkg.in/yaml.v3"
	"os"
)

const (
	FilePath        = "./config/config.yml"
	DefaultHttpPort = 8080
)

type AppConfig struct {
	Redis  RedisConfig
	Server ServerConfig
}

type RedisConfig struct {
	Host     string
	Password string
	Port     uint
}

type ServerConfig struct {
	Port uint
}

func NewAppConfig() (*AppConfig, error) {
	config := &AppConfig{}
	config.setDefaultValues()

	file, err := os.Open(FilePath)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	d := yaml.NewDecoder(file)

	if err = d.Decode(&config); err != nil {
		return nil, err
	}

	return config, nil
}

func (c *AppConfig) setDefaultValues() {
	c.Server.Port = DefaultHttpPort
}
