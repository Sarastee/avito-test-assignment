package repository

import "errors"

const (
	errMsgUserAlreadyRegistered = "user already registered"

	errMsgNoRowsAffected      = "no rows were affected"
	errMsgTagsUniqueViolation = "provided feature and tags pair already exists"
	errMsgBannerNotFound      = "banner not found by provided feature id and tag"
)

var (
	ErrUserAlreadyRegistered = errors.New(errMsgUserAlreadyRegistered) // ErrUserAlreadyRegistered is User Already Registered Error object

	ErrNoRowsAffected      = errors.New(errMsgNoRowsAffected)      // ErrNoRowsAffected is No Rows Affected Error object
	ErrTagsUniqueViolation = errors.New(errMsgTagsUniqueViolation) // ErrTagsUniqueViolation is Tags Unique Violation Error object
	ErrBannerNotFound      = errors.New(errMsgBannerNotFound)      // ErrBannerNotFound is Banner Not Found Error object
)
