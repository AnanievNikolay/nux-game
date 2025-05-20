package domain

import (
	"errors"
)

var (
	ErrorUserIsNil                      = errors.New("user is nil")
	ErrorUsernameWithThisPhoneNotUnique = errors.New(
		"User with this username and phone already exist",
	)
)
