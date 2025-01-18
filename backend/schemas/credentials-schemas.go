package schemas

// LoginCredentials represents the structure for user login credentials.
// It contains the username and password fields which are used for authentication.
type LoginCredentials struct {
	Username string `form:"username" json:"username"` // The username of the user
	Password string `form:"password" json:"password"` // The password of the user
}

// RegisterCredentials represents the credentials required for user registration.
// It includes the user's email, username, and password.
type RegisterCredentials struct {
	Email    string `form:"email"    json:"email"`    // The email of the user
	Username string `form:"username" json:"username"` // The username of the user
	Password string `form:"password" json:"password"` // The password of the user
}

// AuthenticationURL represents the structure for storing the authentication URL.
// URL is the authentication URL for the user, serialized as "authentication_url" in JSON.
type AuthenticationURL struct {
	URL string `json:"authentication_url"` // The authentication URL for the user
}

// JWT represents a JSON Web Token (JWT) used for authentication and authorization.
// It contains a single field, Token, which holds the JWT token as a string.
type JWT struct {
	Token string `json:"token"` // The JWT token
}

// CodeCredentials represents the credentials required for code-based authentication.
// It contains the following fields:
// - Code: A string representing the authentication code. This field is required and is expected to be provided in both form and JSON formats.
type CodeCredentials struct {
	Code string `form:"code" json:"code" binding:"required"` // The code for authentication
}

// TokenCredentials represents the credentials required for token-based authentication.
// It includes the token, username, and email of the user.
// Fields:
// - Token: The token for authentication (required).
// - Username: The username of the user (required).
// - Email: The email of the user (required).
type TokenCredentials struct {
	Token    string `form:"token"    json:"token"    binding:"required"` // The token for authentication
	Username string `form:"username" json:"username" binding:"required"` // The username of the user
	Email    string `form:"email"    json:"email"    binding:"required"` // The email of the user
}
