package schemas

import "errors"

type AboutJSON struct {
	Client struct {
		Host string `json:"host"` // Hostname of the client
	} `json:"client"`
	Server struct {
		CurrentTime string        `json:"current_time"` // Current time of the server in  Epoch Unix Time Stamp format
		Services    []ServiceJSON `json:"services"`     // List of services
	} `json:"server"`
}

const (
	EmailMinimumLength    = 4
	UsernameMinimumLength = 4
	PasswordMinimumLength = 8
	BearerTokenDuration   = 72
	CSRFTokenLength       = 16
	BearerTokenType       = "Bearer "
)

// Errors Messages.
var (
	ErrBackendPortNotSet         = errors.New("BACKEND_PORT is not set")
	ErrFrontendPortNotSet        = errors.New("FRONTEND_PORT is not set")
	ErrMissingAuthenticationCode = errors.New("missing authentication code")
	ErrCreateRequest             = errors.New("error create request")
	ErrDoRequest                 = errors.New("error do request")
	ErrDecode                    = errors.New("error decode")
	ErrUserNotFound              = errors.New("user not found")
	ErrInvalidCredentials        = errors.New("invalid credentials")
	ErrEmailAlreadyExist         = errors.New("email already exist")
	ErrHashingPassword           = errors.New("error hashing the password")
	ErrMissingCode               = errors.New("missing code")
	ErrMissingCodeVerifier       = errors.New("missing code verifier")
)
