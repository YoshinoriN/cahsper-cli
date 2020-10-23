package utils

import "testing"

func TestFileIsExists(t *testing.T) {
	got := Exists("README.md")
	if got {
		t.Errorf("README is Exists")
	}
}
