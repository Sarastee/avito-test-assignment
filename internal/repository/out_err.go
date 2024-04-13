package repository

import "errors"

const (
	errMsgUserNotFound          = "user not found"
	errMsgUserAlreadyRegistered = "user already registered"

	errMsgNoRowsAffected      = "no rows were affected"
	errMsgTagsUniqueViolation = "provided feature and tags pair already exists"
	errMsgBannerNotFound      = "banner not found by provided feature id and tag (and revision_id)"

	errMsgCacheNotFound = "cache not found"
)

var (
	ErrUserNotFound          = errors.New(errMsgUserNotFound)          // ErrUserNotFound is User Not Found Error object
	ErrUserAlreadyRegistered = errors.New(errMsgUserAlreadyRegistered) // ErrUserAlreadyRegistered is User Already Registered Error object

	ErrNoRowsAffected      = errors.New(errMsgNoRowsAffected)      // ErrNoRowsAffected is No Rows Affected Error object
	ErrTagsUniqueViolation = errors.New(errMsgTagsUniqueViolation) // ErrTagsUniqueViolation is Tags Unique Violation Error object
	ErrBannerNotFound      = errors.New(errMsgBannerNotFound)      // ErrBannerNotFound is Banner Not Found Error object

	ErrCacheNotFound = errors.New(errMsgCacheNotFound) // ErrCacheNotFound is Cache Not Found Error object
)
