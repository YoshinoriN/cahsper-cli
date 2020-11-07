package utils

import (
	"bufio"
	"fmt"
	"log"
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
		log.Fatal(err)
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
	secret, err := keyring.Get(createKeyringServiceString(Account), user)
	if err != nil {
		return user, "", err
	}
	return user, secret, err
}

func InteractInputHelper(fieldName string, key CredentialKeyName, userName string, currentValue string) {
	fmt.Printf("%s [%s]: ", fieldName, currentValue)

	scanner := bufio.NewScanner(os.Stdin)
	var newValue string
	for scanner.Scan() {
		newValue = scanner.Text()
		break
	}

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
