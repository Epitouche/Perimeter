package schemas

import "errors"

// Errors Messages.
var (
	ErrBackendPortNotSet         = errors.New("BACKEND_PORT is not set")
	ErrMissingAuthenticationCode = errors.New("missing authentication code")
)
