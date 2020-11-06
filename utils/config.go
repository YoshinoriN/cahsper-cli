package utils

import (
	"fmt"
	"os"
	"path/filepath"

	"gopkg.in/yaml.v2"
)

func GetConfigFileDir() string {
	homeDir := GetUserHomeDirectory()
	return filepath.Join(homeDir, ".cahsper")
}

func GetConfigFilePath() string {
	return filepath.Join(GetConfigFileDir(), ".config")
}

type Config struct {
	Settings struct {
		Cognito struct {
			UserPoolID  string `yaml:"userPoolID"`
			AppClientID string `yaml:"appClientId"`
		}
		ServerURL string `yaml:"serverUrl"`
	}
}

func Read(configFilePath string) Config {
	file, err := os.Open(configFilePath)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	d := yaml.NewDecoder(file)
	var config Config
	if err := d.Decode(&config); err != nil {
		panic(err)
	}
	return config
}

func Print(config Config) {
	// TODO
	fmt.Println(config)
}
