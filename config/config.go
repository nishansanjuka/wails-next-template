package config

import (
	"encoding/json"
	"os"
)

type Config struct {
	SomeConfig string `json:"config_data"`
}

func LoadConfig() *Config {
	file, _ := os.Open("config/config.json")

	defer file.Close()
	var config Config
	json.NewDecoder(file).Decode(&config)

	return &config
}
