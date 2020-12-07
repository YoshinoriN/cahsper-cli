package utils

import (
	"os"
	"os/user"
)

// GetUserHomeDirectory get user home firectory
func GetUserHomeDirectory() string {
	user, err := user.Current()
	if err != nil {
		panic(err)
	}
	return user.HomeDir
}

// MakeDir make directory
func MakeDir(path string) {
	if _, err := os.Stat(path); os.IsNotExist(err) {
		os.Mkdir(path, 600)
	}
}
