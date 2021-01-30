package service

import "errors"

// ErrMustHaveUsername ...
var ErrMustHaveUsername = errors.New("'Username' field is required")

// ErrMustHaveEmail ...
var ErrMustHaveEmail = errors.New("'Email' field is required")

// ErrInvalidPassword ...
var ErrInvalidPassword = errors.New("Password is invalid")
