package config

import (
	"os"

	"gopkg.in/yaml.v2"
)

type Config struct {
	Server struct {
		Port string `yaml:"port" env:"PORT" env-default:"8080"`
		Host string `yaml:"host" env:"PORT" env-default:"8080"`
	} `yaml:"server"`

	Database struct {
		Username string `yaml:"user"`
		Password string `yaml:"pass"`
	} `yaml:"database"`

	JWTSecretKey string `yaml:"secret_key"`
}

func NewConfig(filePath string) (*Config, error) {
	f, err := os.Open(filePath)
	if err != nil {
		return nil, err
	}
	defer f.Close()

	var cfg *Config
	decoder := yaml.NewDecoder(f)
	err = decoder.Decode(&cfg)
	if err != nil {
		return nil, err
	}
	return cfg, nil
}
