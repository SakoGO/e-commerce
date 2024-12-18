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

	Database struct {
		User      string `json:"user"`
		Password  string `json:"password"`
		Host      string `json:"host"`
		Port      int    `json:"port"`
		Name      string `json:"name"`
		Charset   string `json:"charset"`
		ParseTime bool   `json:"parse_time"`
		Loc       string `json:"loc"`
	} `json:"database"`
}

func LoadConfig(filePath string) (*Config, error) {

	// Open config file
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
