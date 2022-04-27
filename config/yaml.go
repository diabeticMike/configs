package config

import (
	"gopkg.in/yaml.v3"
	"os"
)

func NewYaml(filePath string) (Config, error) {
	file, err := os.Open(filePath)
	defer file.Close()
	if err != nil {
		return Config{}, err
	}

	var config Config
	err = yaml.NewDecoder(file).Decode(&config)
	if err != nil {
		return Config{}, err
	}

	return config, nil
}
