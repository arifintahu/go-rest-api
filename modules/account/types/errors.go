package types

import "errors"

var (
	ErrAccountEmailNotFound = errors.New("Account email is not found")
	ErrAccountPasswordInvalid = errors.New("Password invalid")
)
