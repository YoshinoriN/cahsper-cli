package utils

import (
	"os"
)

// Exists return file is exists or not
func Exists(path string) bool {
	_, err := os.Stat(path)
	return err == nil
}
