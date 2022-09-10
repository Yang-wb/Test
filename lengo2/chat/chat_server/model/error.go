package model

import "errors"

var (
	ErrUserNotExist    = errors.New("user not exist")
	ErrInvalidPassword = errors.New("Passwd or username not right")
	ErrInvalidParams   = errors.New("Invalid params")
	ErrUserExist       = errors.New("user exist")
)
