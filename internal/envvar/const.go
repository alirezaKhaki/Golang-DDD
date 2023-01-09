package envvar

import "os"

// GetSecret returns the jwt secret.
func AuthSecret() string {
	secret, exists := os.LookupEnv("AUTH_SECRET")
	if !exists {
		secret = "secret_key"
	}
	return secret
}