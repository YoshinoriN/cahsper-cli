package utils

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"

	"gopkg.in/yaml.v2"
)

// GetConfigFileDir get config file dir
func GetConfigFileDir() string {
	homeDir := GetUserHomeDirectory()
	return filepath.Join(homeDir, ".cahsper")
}

// GetConfigFilePath get config file path
func GetConfigFilePath() string {
	return filepath.Join(GetConfigFileDir(), ".config")
}

// Config structure for config file
type Config struct {
	Settings struct {
		Aws struct {
			Region  string `yaml:"region"`
			Cognito struct {
				UserPoolID  string `yaml:"userPoolID"`
				AppClientID string `yaml:"appClientId"`
			} `yaml:"cognito"`
		} `yaml:"aws"`
		ServerURL string `yaml:"serverUrl"`
	} `yaml:"settings"`
}

// Read read config values from config file
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

// Write write config values to config file
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

// Print print config values from config file
func Print(config Config) {
	fmt.Print("settings:\n")
	fmt.Print("  aws:\n")
	fmt.Printf("    region: %s\n", config.Settings.Aws.Region)
	fmt.Print("    cognito:\n")
	fmt.Printf("      userPoolID: %s\n", config.Settings.Aws.Cognito.UserPoolID)
	fmt.Printf("      appClientId: %s\n", config.Settings.Aws.Cognito.AppClientID)
	fmt.Printf("  serverUrl: %s", config.Settings.ServerURL)
}
