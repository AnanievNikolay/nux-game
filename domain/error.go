package domain

import (
	"errors"
)

var (
	ErrorUserNotFound                   = errors.New("user not found")
	ErrorUsernameWithThisPhoneNotUnique = errors.New(
		"user with this username and phone already exist",
	)
	ErrTokenInvalidOrExpired = errors.New("token invalid or expired")
)
