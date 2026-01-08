package auth

import "errors"

// replace with real JWT validation
func ValidateJWT(token string) (string, error) {
	if token == "" {
		return "", errors.New("empty token")
	}

	// TODO:
	// - verify signature
	// - check expiration
	// - parse claims

	// example user id
	return "user-123", nil
}
