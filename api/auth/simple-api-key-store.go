package auth

import "errors"

// SimpleAPIKeyStore with a fixed key
type SimpleAPIKeyStore struct {
	apiKey string
}

// NewSimpleAPIKeyStore creates a new instance
func NewSimpleAPIKeyStore(key string) (*SimpleAPIKeyStore, error) {
	if key == "" {
		return nil, errors.New("key must be provided")
	}

	return &SimpleAPIKeyStore{
		apiKey: key,
	}, nil
}

// GetAPIKey from the store
func (s *SimpleAPIKeyStore) GetAPIKey() string {
	return s.apiKey
}
