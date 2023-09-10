package utils

import (
	"errors"
)

var (
	ErrTokenSigningMethod = errors.New("ErrTokenSigningMethod")
	ErrTokenExpired       = errors.New("ErrTokenExpired")
	ErrTokenMalformed     = errors.New("ErrTokenMalformed")
	ErrTokenHandle        = errors.New("ErrTokenHandle")
	ErrSignedString       = errors.New("ErrSignedString")
)
