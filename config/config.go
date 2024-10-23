package config

import (
	"os"
)

type Config struct {
	
}

func NewConfig(configFile string) (*Config, error) {
	file, err := os.Open(configFile)
	if err != nil {
		return nil, err
	}
	defer file.Close()
	cfg := &Config{}
	


	if err != nil {
		return nil, err
	}
	return cfg, nil
}
