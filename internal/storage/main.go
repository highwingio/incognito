package storage

import (
	"github.com/99designs/keyring"
	"github.com/highwingio/incognito/internal/types"
)

const usernameKey = "Username"
const passwordKey = "Password"
const userPoolIDKey = "UserPoolID"
const clientIDKey = "ClientID"
const serviceName = "io.highwing.incognito"

// KeyringStorage manages storage of credentials via Keyring
type KeyringStorage struct {
	Keyring keyring.Keyring
}

// NewKeyringStorage initializes and returns a new instance of a KeyringStorage
func NewKeyringStorage() (*KeyringStorage, error) {
	ring, err := keyring.Open(keyring.Config{
		ServiceName: serviceName,
	})

	if err != nil {
		return nil, err
	}

	return &KeyringStorage{
		Keyring: ring,
	}, nil
}

// StoreLoginCredentials attempts to save the provided credentials via Keyring
func (k *KeyringStorage) StoreLoginCredentials(
	credentials *types.CognitoAuthenticationCredentials,
) error {
	for key, value := range map[string]string{
		usernameKey:   credentials.Username,
		passwordKey:   credentials.Password,
		userPoolIDKey: credentials.UserPoolID,
		clientIDKey:   credentials.ClientID,
	} {
		err := k.Keyring.Set(keyring.Item{
			Key:  key,
			Data: []byte(value),
		})
		if err != nil {
			return err
		}
	}
	return nil
}

// RetrieveLoginCredentials attempts to retrieve credentials from Keyring
func (k *KeyringStorage) RetrieveLoginCredentials() (*types.CognitoAuthenticationCredentials, error) {
	username, err := k.Keyring.Get(usernameKey)
	if err != nil {
		return nil, err
	}

	password, err := k.Keyring.Get(passwordKey)
	if err != nil {
		return nil, err
	}

	userPoolID, err := k.Keyring.Get(userPoolIDKey)
	if err != nil {
		return nil, err
	}

	clientID, err := k.Keyring.Get(clientIDKey)
	if err != nil {
		return nil, err
	}

	return &types.CognitoAuthenticationCredentials{
		Username:   string(username.Data),
		Password:   string(password.Data),
		UserPoolID: string(userPoolID.Data),
		ClientID:   string(clientID.Data),
	}, nil
}
