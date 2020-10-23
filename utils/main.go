package utils

import (
	"os"
	"os/user"
)

func GetUserHomeDirectory() string {
	user, err := user.Current()
	if err != nil {
		panic(err)
	}
	return user.HomeDir
}

func MakeDir(path string) {
	if _, err := os.Stat(path); os.IsNotExist(err) {
		os.Mkdir(path, 600)
	}
}
