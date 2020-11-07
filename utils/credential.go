package utils

import (
	"log"
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
