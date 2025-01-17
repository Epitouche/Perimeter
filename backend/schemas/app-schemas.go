package schemas

import "errors"

// AboutJSON represents the structure of the JSON response containing information
// about the client and server. It includes the client's hostname and the server's
// current time and list of services.
type AboutJSON struct {
	Client struct {
		Host string `json:"host"` // Hostname of the client
	} `json:"client"` // Client information
	Server struct {
		CurrentTime string        `json:"current_time"` // Current time of the server in  Epoch Unix Time Stamp format
		Services    []ServiceJSON `json:"services"`     // List of services
	} `json:"server"` // Server information
}

const (
	EmailMinimumLength    = 4         // Minimum length of an email address
	UsernameMinimumLength = 4         // Minimum length of a username
	PasswordMinimumLength = 8         // Minimum length of a password
	BearerTokenDuration   = 72        // Duration of the bearer token in hours
	CSRFTokenLength       = 16        // Length of the CSRF token
	BearerTokenType       = "Bearer " // Bearer token type
)

// Errors Messages.
var (
	ErrBackendPortNotSet          = errors.New("BACKEND_PORT is not set")           // Error message for missing BACKEND_PORT environment variable
	ErrFrontendPortNotSet         = errors.New("FRONTEND_PORT is not set")          // Error message for missing FRONTEND_PORT environment variable
	ErrFrontendExternalHostNotSet = errors.New("FRONTEND_EXTERNAL_HOST is not set") // Error message for missing FRONTEND_EXTERNAL_HOST environment variable
	ErrIsProductionNotSet         = errors.New("IS_PRODUCTION is not set")          // Error message for missing IS_PRODUCTION environment variable
	ErrMissingAuthenticationCode  = errors.New("missing authentication code")       // Error message for missing authentication code
	ErrCreateRequest              = errors.New("error create request")              // Error message for failed request creation
	ErrDoRequest                  = errors.New("error do request")                  // Error message for failed request execution
	ErrDecode                     = errors.New("error decode")                      // Error message for failed decoding
	ErrUserNotFound               = errors.New("user not found")                    // Error message for user not found
	ErrInvalidCredentials         = errors.New("invalid credentials")               // Error message for invalid credentials
	ErrEmailAlreadyExist          = errors.New("email already exist")               // Error message for existing email
	ErrInvalidEmail               = errors.New("invalid email")                     // Error message for invalid email
	ErrHashingPassword            = errors.New("error hashing the password")        // Error message for hashing password
	ErrMissingCode                = errors.New("missing code")                      // Error message for missing code
	ErrMissingCodeVerifier        = errors.New("missing code verifier")             // Error message for missing code verifier
)
