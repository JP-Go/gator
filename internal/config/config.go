package config

import (
	"bytes"
	"encoding/json"
	"os"
	"path"
)

var configFilename = ".gatorconfig.json"

func getConfigFilePath(filename string) (string, error) {
	homePath, err := os.UserHomeDir()
	if err != nil {
		return "", err
	}
	configFilePath := path.Join(homePath, filename)
	return configFilePath, nil
}

func Read() (*Config, error) {
	filepath, err := getConfigFilePath(configFilename)
	if err != nil {
		return nil, err
	}
	contents, err := os.ReadFile(filepath)
	if err != nil {
		return nil, err
	}
	config, err := readFromBytes(contents)
	if err != nil {
		return nil, err
	}
	return &config, nil
}

func readFromBytes(contents []byte) (Config, error) {
	var config Config
	if err := json.NewDecoder(bytes.NewReader(contents)).Decode(&config); err != nil {
		return Config{}, err
	}
	return config, nil
}

func write(cfg Config) error {
	fullpath, err := getConfigFilePath(configFilename)
	if err != nil {
		return err
	}
	file, err := os.Create(fullpath)
	if err != nil {
		return err
	}
	defer file.Close()
	newContents, err := json.Marshal(cfg)
	if err != nil {
		return err
	}
	if _, err := file.Write(newContents); err != nil {
		return err
	}
	return nil
}

func (cfg *Config) SetUser(username string) error {
	cfg.CurrentUserName = username
	err := write(*cfg)
	return err
}
