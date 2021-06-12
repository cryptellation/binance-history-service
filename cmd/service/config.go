package main

import "github.com/pelletier/go-toml"

// Config is the configuration for the messenger binance server
type Config struct {
	APIKey    string `toml:"api_key"`
	SecretKey string `toml:"secret_key"`
}

// ConfigFromFile will generate a configuration structure from a TOML config file
func ConfigFromFile(path string) (c Config, err error) {
	// Load file
	tree, err := toml.LoadFile(path)
	if err != nil {
		return Config{}, err
	}

	// Change into structure
	err = tree.Unmarshal(&c)
	return c, err
}
