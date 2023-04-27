package config 

import (
	"io/ioutil"
	"log"

	"gopkg.in/yaml.v3"
)

// Config represents the application configuration.
type Config struct {
	Port         int    `yaml:"port"`
	DataPath     string `yaml:"data_path"`
	InitialNode  string `yaml:"initial_node"`
	InitialKey   string `yaml:"initial_key"`
}

// LoadConfig loads the application configuration from the given file path.
func LoadConfig(path string) (*Config, error) {
	data, err := ioutil.ReadFile(path)
	if err != nil {
		log.Fatalf("error reading config file: %v", err)
		return nil, err
	}

	var cfg Config 
	err = yaml.Unmarshal(data, &cfg)
	if err != nil {
		log.Fatalf("error parsing config file: %v", err)
		return nil, err
	}

	return &cfg, nil
}

