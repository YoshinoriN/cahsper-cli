package utils

import (
	"log"
	"testing"

	"github.com/zalando/go-keyring"
)

const userName = "Jhon Doe"
const pass = "password!!"

func TestShouldSetAccountCredential(t *testing.T) {
	err := SetCredential(userName, Account, pass)
	if err != nil {
		t.Errorf("Can not set credential")
	}
}

func TestShouldGetAccountCredential(t *testing.T) {
	u, p, err := GetAccount(userName)
	if err != nil {
		t.Errorf("Can not set credential")
	}
	if u != userName {
		t.Errorf("Can not get user name")
	}

	if p != "password!!" {
		t.Errorf("Can not get password")
	}

}

// TODO: Is there exists better way?
func TestTearDown(t *testing.T) {
	err := keyring.Delete(createKeyringServiceString(Account), userName)
	if err != nil {
		log.Fatal(err)
	}
}
