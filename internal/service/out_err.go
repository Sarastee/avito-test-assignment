package service

import "errors"

const (
	errMsgPasswordToLong = "password to long"
)

var (
	ErrPasswordToLong = errors.New(errMsgPasswordToLong) // ErrPasswordToLong is Password To Long Error object
)
