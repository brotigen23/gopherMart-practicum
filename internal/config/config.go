package config

import (
	"flag"

	"github.com/ilyakaznacheev/cleanenv"
)

type Config struct {
	RunAdress            string `env:"RUN_ADDRESS" env-default:"localhost:8080"`
	DatabaseURI          string `env:"DATABASE_URI"`
	AccrualSystemAddress string `env:"ACCRUAL_SYSTEM_ADDRESS"`
	JWTSecretKey         string `env:"secret_key" env-default:"secret"`
}

func NewConfig() (*Config, error) {
	// Read env
	cfg := &Config{}
	cleanenv.ReadEnv(cfg)
	// Read flags
	a := flag.String("a", "", "server address")          // RUN_ADDRESS
	d := flag.String("d", "", "database connect string") // DATABASE_URI
	r := flag.String("r", "", "accrual system address")  // ACCRUAL_SYSTEM_ADDRESS

	flag.Parse()
	if *a != "" {
		cfg.RunAdress = *a
	}
	if *d != "" {
		cfg.DatabaseURI = *d
	}
	if *r != "" {
		cfg.AccrualSystemAddress = *r
	}

	return cfg, nil
}
