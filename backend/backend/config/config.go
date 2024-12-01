package config

import (
	"encoding/json"
	"fmt"
	"os"
)

type Config struct {
	Server struct {
		Addr         string `json:"addr"`
		ReadTimeout  int    `json:"readTimeout"`
		WriteTimeout int    `json:"WriteTimeout"`
		IdleTimeout  int    `json:"ShutdownTimeout"`
	} `json:"server"`
}

func LoadConfig(filePath string) (*Config, error) {

	// Open file
	file, err := os.Open(filePath)
	if err != nil {
		return nil, fmt.Errorf("unable to open config file: %v", err)
	}
	defer file.Close()

	// Decode json to config struct
	var cfg Config
	decoder := json.NewDecoder(file)
	err = decoder.Decode(&cfg)
	if err != nil {
		return nil, fmt.Errorf("unable to open config: %v", err)
	}

	return &cfg, nil
}
