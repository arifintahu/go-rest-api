package types

import "errors"

var (
	ErrUserEmailExist = errors.New("User email is exist")
	ErrUserNotFound = errors.New("User is not found")
)
