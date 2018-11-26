package application

import "errors"

var (
	ErrUserNotFound     = errors.New("user not found")
	ErrUserAlreadyExist = errors.New("user already exist")
	ErrPasswordNotMatch = errors.New("password not match")
)
