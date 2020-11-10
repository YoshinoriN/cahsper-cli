package utils

import (
	"fmt"
	"io/ioutil"
	"log"
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
		} `yaml:"cognito"`
		ServerURL string `yaml:"serverUrl"`
	} `yaml:"settings"`
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

func Write(configFilePath string, config Config) {
	d, err := yaml.Marshal(&config)
	if err != nil {
		log.Fatalf("error: %v", err)
	}
	err = ioutil.WriteFile(configFilePath, d, 0644)
	if err != nil {
		log.Fatal(err)
	}
}

func Print(config Config) {
	// TODO
	fmt.Println(config)
}
