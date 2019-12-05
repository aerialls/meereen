package config

import (
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"

	"gopkg.in/yaml.v2"
)

// NewConfig loads the configuration from the YAML file
func NewConfig(path string) (*Config, error) {
	data, err := ioutil.ReadFile(path)
	if err != nil {
		return nil, fmt.Errorf("unable to load the config from file %s", err)
	}

	cfg := &Config{}
	*cfg = DefaultConfig

	err = yaml.Unmarshal([]byte(data), &cfg)
	if err != nil {
		return nil, err
	}

	return cfg, nil
}

// Validate verifis
func (c *Config) Validate() error {
	if _, err := os.Stat(c.ChecksFolder); os.IsNotExist(err) {
		return fmt.Errorf("checks folder %s does not exist", c.ChecksFolder)
	}

	if c.Global.Delta == 0 {
		c.Global.Delta = DefaultConfig.Global.Delta
	}

	return nil
}

// GetChecks returns the config for all checks inside the directory
func (c *Config) GetChecks() ([]CheckConfig, error) {
	checks := []CheckConfig{}
	err := filepath.Walk(c.ChecksFolder, func(path string, info os.FileInfo, err error) error {
		if !info.IsDir() {
			data, err := ioutil.ReadFile(path)
			if err != nil {
				return err
			}

			cfg := ChecksConfig{}
			err = yaml.Unmarshal([]byte(data), &cfg)
			if err != nil {
				return err
			}
			checks = append(checks, cfg.Checks...)
		}

		return nil
	})

	if err != nil {
		return nil, err
	}

	return checks, nil
}
