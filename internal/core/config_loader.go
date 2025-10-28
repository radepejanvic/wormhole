package core

import (
	"fmt"
	"os"

	"github.com/c12s/wormhole/internal/types"
	"gopkg.in/yaml.v3"
)

func LoadFromYaml(path string) (*types.Config, error) {
	data, err := os.ReadFile(path)
	if err != nil {
		return nil, fmt.Errorf("read file: %w", err)
	}

	var conf types.Config
	err = yaml.Unmarshal(data, &conf)
	if err != nil {
		return nil, fmt.Errorf("unmarshal yaml: %w", err)
	}

	err = conf.Validate()
	if err != nil {
		return nil, fmt.Errorf("validate config: %w", err)
	}

	return &conf, nil
}
