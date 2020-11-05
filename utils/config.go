package utils

import (
	"path/filepath"
)

func GetConfigFileDir() string {
	homeDir := GetUserHomeDirectory()
	return filepath.Join(homeDir, ".cahsper")
}

func GetConfigFilePath() string {
	return filepath.Join(GetConfigFileDir(), ".config")
}
