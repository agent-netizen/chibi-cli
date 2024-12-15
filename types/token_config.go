package types

import (
	"encoding/json"
	"os"
	"path"
)

type TokenConfig struct {
	TokenType string `json:"token_type"`
	ExpiresIn int `json:"expires_in"`
	AccessToken string `json:"access_token"`
	RefreshToken string `json:"refresh_token"`
}

func createConfigDir() string {
	osConfigPath, _ := os.UserConfigDir()
	configDir := path.Join(osConfigPath, "chibi")
	_, err := os.Stat(configDir)

	if err == nil {
		os.RemoveAll(configDir)
	}
	os.MkdirAll(configDir, 0755)
	return configDir
}

func (c TokenConfig) FlushToJsonFile() error {
	configDir := createConfigDir()
	jsonPath := path.Join(configDir, "token.json")
	jsonString, err := json.Marshal(c)
	if err != nil {
		return err
	}
	err = os.WriteFile(jsonPath, jsonString, 0644)
	if err != nil {
		return err
	}
	return nil
}

func (c *TokenConfig) ReadFromJsonFile() error {
	osConfigDir, _ := os.UserConfigDir()
	configFilePath := path.Join(osConfigDir, "chibi", "token.json")
	file, err := os.ReadFile(configFilePath)
	if err != nil {
		return err
	}
	err = json.Unmarshal(file, &c)
	if err != nil {
		return err
	}
	return nil
}

func NewTokenConfig() *TokenConfig {
	return &TokenConfig{}
}