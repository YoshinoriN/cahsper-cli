package utils

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"github.com/zalando/go-keyring"
)

// CredentialKeyName type for CredentialKeyName
type CredentialKeyName string

const (
	// ServiceName key name for keyring
	ServiceName CredentialKeyName = "Cahsper"
	// Account key name for keyring
	Account CredentialKeyName = "Account"
	// IDToken key name for keyring
	IDToken CredentialKeyName = "IdToken"
	// AccessToken key name for keyring
	AccessToken CredentialKeyName = "AccessToken"
	// RefreshToken key name for keyring
	RefreshToken CredentialKeyName = "RefreshToken"
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
	case RefreshToken:
		return "RefreshToken"
	default:
		return "" // TODO: error handling
	}
}

func createKeyringServiceString(key CredentialKeyName) string {
	return strings.Join([]string{ServiceName.String(), key.String()}, ":")
}

// SetCredential set credential to keyring
func SetCredential(user string, key CredentialKeyName, value string) error {
	err := keyring.Set(createKeyringServiceString(key), user, value)
	if err != nil {
		return err
	}
	return nil
}

// GetCredential get credential from keyring
func GetCredential(user string, key CredentialKeyName) (string, error) {
	secret, err := keyring.Get(createKeyringServiceString(key), user)
	if err != nil {
		return "", err
	}
	return secret, err
}

// GetAccount get account from keyring
func GetAccount(user string) (string, string, error) {
	account, err := keyring.Get(createKeyringServiceString(Account), user)
	if err != nil {
		return user, "", err
	}
	return user, account, err
}

// InteractInputHelper helper function for input credential info
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
