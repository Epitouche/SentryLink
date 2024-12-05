package tools

import (
	"crypto/rand"
	"encoding/hex"
)

type CSRF interface {
	GenerateCSRFToken() (string, error)
}

// Generate a random CSRF token.
func GenerateCSRFToken() (string, error) {
	bytes := make([]byte, 16)
	_, err := rand.Read(bytes)
	if err != nil {
		return "", err
	}
	return hex.EncodeToString(bytes), nil
}
