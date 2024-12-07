package schemas

import "errors"

// Errors Messages.
var (
	ErrBackendPortNotSet         = errors.New("BACKEND_PORT is not set")
	ErrMissingAuthenticationCode = errors.New("missing authentication code")
	ErrCreateRequest             = errors.New("error create request")
	ErrDoRequest                 = errors.New("error do request")
	ErrDecode                    = errors.New("error decode")
	ErrUserNotFound              = errors.New("user not found")
	ErrInvalidCredentials        = errors.New("invalid credentials")
	ErrEmailAlreadyExist         = errors.New("email already exist")
	ErrHashingPassword           = errors.New("error hashing the password")
)
