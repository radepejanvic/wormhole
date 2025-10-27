package core

import (
	"fmt"
	"os"

	"gopkg.in/yaml.v3"
)

type Config struct {
	NodeCount   uint8  `yaml:"nodeCount"`
	OSDistro    string `yaml:"osDistro"`
	BaseAddress string `yaml:"baseAddress"`
	BasePort    uint16 `yaml:"basePort"`
}

func NewFromYaml(path string) (*Config, error) {
	data, err := os.ReadFile(path)
	if err != nil {
		return nil, fmt.Errorf("error reading file: %w", err)
	}

	var conf Config
	err = yaml.Unmarshal(data, &conf)
	if err != nil {
		return nil, fmt.Errorf("error unmarshaling yaml: %w", err)
	}

	return &conf, nil
}
