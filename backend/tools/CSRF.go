package tools

import (
	"crypto/rand"
	"encoding/hex"
	"fmt"

	"github.com/Epitouche/Perimeter/schemas"
)

type CSRF interface {
	GenerateCSRFToken() (string, error)
}

// Generate a random CSRF token.
func GenerateCSRFToken() (string, error) {
	bytes := make([]byte, schemas.CSRFTokenLength)
	_, err := rand.Read(bytes)
	if err != nil {
		return "", fmt.Errorf("unable to generate CSRF token: %w", err)
	}
	return hex.EncodeToString(bytes), nil
}
