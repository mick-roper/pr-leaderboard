package auth

type (
	// APIKeyStore that can be used to retrieve an API key
	APIKeyStore interface {
		GetAPIKey() string
	}
)
