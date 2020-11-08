package utils

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"github.com/zalando/go-keyring"
)

type CredentialKeyName string

const (
	ServiceName CredentialKeyName = "Cahsper"
	Account     CredentialKeyName = "Account"
	IDToken     CredentialKeyName = "IdToken"
	AccessToken CredentialKeyName = "AccessToken"
)

func (credentialKeyName CredentialKeyName) String() string {
	switch credentialKeyName {
	case ServiceName:
		return "Cahsper"
	case Account:
		return "Account"
	case IDToken:
		return "IdToken"
	case AccessToken:
		return "AccessToken"
	default:
		return "" // TODO: error handling
	}
}

func createKeyringServiceString(key CredentialKeyName) string {
	return strings.Join([]string{ServiceName.String(), key.String()}, ":")
}

func SetCredential(user string, key CredentialKeyName, value string) error {
	err := keyring.Set(createKeyringServiceString(key), user, value)
	if err != nil {
		return err
	}
	return nil
}

func GetCredential(user string, key CredentialKeyName) (string, error) {
	secret, err := keyring.Get(createKeyringServiceString(key), user)
	if err != nil {
		return "", err
	}
	return secret, err
}

func GetAccount(user string) (string, string, error) {
	account, err := keyring.Get(createKeyringServiceString(Account), user)
	if err != nil {
		return user, "", err
	}
	return user, account, err
}

func InteractInputHelper(fieldName string, key CredentialKeyName, userName string, currentValue string) {
	if strings.TrimSpace(currentValue) != "" {
		fmt.Printf("%s [%s]: ", fieldName, currentValue)
	} else {
		fmt.Printf("%s: ", fieldName)
	}

	scanner := bufio.NewScanner(os.Stdin)
	var newValue string
	scanner.Scan()
	newValue = scanner.Text()

	if strings.TrimSpace(newValue) == "" {
		if strings.TrimSpace(currentValue) == "" {
			fmt.Printf("%s required", fieldName)
			os.Exit(0)
		}
		newValue = currentValue
	}
	err := SetCredential(userName, key, newValue)
	if err != nil {
		os.Exit(1)
	}
}