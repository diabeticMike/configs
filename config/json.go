package config

import (
	"encoding/json"
	"os"
)

type Config struct {
	ListeningUrl string `json:"ListeningUrl" yaml:"ListeningUrl" envconfig:"LISTENING_URL"`
	Delay        uint   `json:"Delay" yaml:"Delay" envconfig:"DELAY"`
	Inner        Inner  `json:"Inner"`
}

type Inner struct {
	Field string `json:"Field"`
}

func NewJson(filePath string) (Config, error) {
	file, err := os.Open(filePath)
	defer file.Close()
	if err != nil {
		return Config{}, err
	}

	var config Config
	err = json.NewDecoder(file).Decode(&config)
	if err != nil {
		return Config{}, err
	}

	return config, nil
}
