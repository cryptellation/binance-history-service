package main

import "testing"

func defaultConfig() (Config, error) {
	return ConfigFromFile("../../configs/binance-history-service.example.toml")
}

func TestConfigFromFile(t *testing.T) {
	c, err := defaultConfig()
	if err != nil {
		t.Fatal("Cannot get default config:", err)
	}

	// Test APIKey
	if c.APIKey != "your api key" {
		t.Error("Default API key is incorrect:", c.APIKey)
	}

	// Test SecretKey
	if c.SecretKey != "your secret key" {
		t.Error("Default Secret key is incorrect:", c.SecretKey)
	}
}

func TestConfigFromFile_NoFile(t *testing.T) {
	if _, err := ConfigFromFile("/do/not/exist"); err == nil {
		t.Fatal("There should be no file")
	}
}
