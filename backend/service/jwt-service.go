package service

import (
	"errors"
	"fmt"
	"os"
	"strconv"
	"time"

	"github.com/golang-jwt/jwt"

	"area/schemas"
)

// JWTService defines the methods that any implementation of JWT authentication should provide.
// It includes methods for generating, validating, and extracting user information from JWT tokens.
type JWTService interface {
	// GenerateToken generates a JWT token for a given user.
	// Parameters:
	// - userID: The unique identifier of the user.
	// - name: The name of the user.
	// - admin: A boolean indicating if the user has admin privileges.
	// Returns:
	// - A JWT token as a string.
	GenerateToken(userID string, name string, admin bool) string

	// ValidateToken validates a given JWT token.
	// Parameters:
	// - tokenString: The JWT token as a string.
	// Returns:
	// - A pointer to a jwt.Token if the token is valid.
	// - An error if the token is invalid.
	ValidateToken(tokenString string) (*jwt.Token, error)

	// GetUserIdfromJWTToken extracts the user ID from a given JWT token.
	// Parameters:
	// - tokenString: The JWT token as a string.
	// Returns:
	// - The user ID as a uint64.
	// - An error if the extraction fails.
	GetUserIdfromJWTToken(tokenString string) (userID uint64, err error)
}

// jwtCustomClaims are custom claims extending default ones.
type jwtCustomClaims struct {
	Name  string `json:"name"`
	Admin bool   `json:"admin"`
	jwt.StandardClaims
}

// jwtService represents a service for handling JWT (JSON Web Token) operations.
// It contains the secret key used for signing tokens and the issuer of the tokens.
type jwtService struct {
	secretKey string
	issuer    string
}

// NewJWTService creates a new instance of JWTService with a secret key and issuer.
// The secret key is retrieved using the getSecretKey function, and the issuer is set to "email@example.com".
// Returns a JWTService interface.
func NewJWTService() JWTService {
	return &jwtService{
		secretKey: getSecretKey(),
		issuer:    "email@example.com",
	}
}

// getSecretKey retrieves the JWT secret key from the environment variable "JWT_SECRET".
// If the environment variable is not set, it will panic with an appropriate error message.
// Returns the secret key as a string.
func getSecretKey() string {
	secret := os.Getenv("JWT_SECRET")
	if secret == "" {
		panic("JWT_SECRET is not set")
	}
	return secret
}

func (jwtSrv *jwtService) GenerateToken(userID string, username string, admin bool) string {
	// Set custom and standard claims
	claims := &jwtCustomClaims{
		username,
		admin,
		jwt.StandardClaims{
			ExpiresAt: time.Now().Add(time.Hour * schemas.BearerTokenDuration).Unix(),
			Issuer:    jwtSrv.issuer,
			IssuedAt:  time.Now().Unix(),
			Id:        userID,
		},
	}

	// Create token with claims
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	// Generate encoded token using the secret signing key
	t, err := token.SignedString([]byte(jwtSrv.secretKey))
	if err != nil {
		panic(err)
	}
	return t
}

// ValidateToken validates a given JWT token string and returns the parsed token if valid.
// It checks the signing method and uses the secret key to validate the token's signature.
//
// Parameters:
//   - tokenString: The JWT token string to be validated.
//
// Returns:
//   - *jwt.Token: The parsed JWT token if validation is successful.
//   - error: An error if the token is invalid or if there is an issue during parsing.
func (jwtSrv *jwtService) ValidateToken(tokenString string) (*jwt.Token, error) {
	result, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		// Signing method validation
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		// Return the secret signing key
		return []byte(jwtSrv.secretKey), nil
	})
	return result, err
}

// GetUserIdfromJWTToken extracts the user ID from a given JWT token string.
// It validates the token and parses the "jti" claim to retrieve the user ID.
//
// Parameters:
//   - tokenString: The JWT token string from which to extract the user ID.
//
// Returns:
//   - userID: The extracted user ID as a uint64.
//   - err: An error if the token is invalid, the "jti" claim is missing or not a string,
//     or if the "jti" claim cannot be parsed as a uint64.
func (jwtSrv *jwtService) GetUserIdfromJWTToken(tokenString string) (userID uint64, err error) {
	token, err := jwtSrv.ValidateToken(tokenString)
	if err != nil {
		return 0, err
	}

	if token.Valid {
		claims := token.Claims.(jwt.MapClaims)
		if jti, ok := claims["jti"].(string); ok {
			id, err := strconv.ParseUint(jti, 10, 64)
			if err != nil {
				return 0, errors.New("jti claim is not a float64")
			}
			return id, nil
		}
		return 0, errors.New("jti claim is not a float64")
	} else {
		return 0, errors.New("invalid token")
	}
}
