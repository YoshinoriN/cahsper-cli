package utils

import "testing"

const userName = "Jhon Doe"
const pass = "password!!"

func TestShouldSetAccountCredential(t *testing.T) {
	err := SetCredencial(userName, Account, pass)
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
