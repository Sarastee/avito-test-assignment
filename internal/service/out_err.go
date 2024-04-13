package service

import "errors"

const (
	errMsgPasswordToLong = "password to long"
	errMsgWrongPassword  = "wrong password"
)

var (
	ErrPasswordToLong = errors.New(errMsgPasswordToLong) // ErrPasswordToLong is Password To Long Error object
	ErrWrongPassword  = errors.New(errMsgWrongPassword)  // ErrWrongPassword is Wrong Password Error object
)
