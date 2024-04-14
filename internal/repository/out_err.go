package repository

import "errors"

const (
	errMsgUserNotFound          = "user not found"
	errMsgUserAlreadyRegistered = "user already registered"

	errMsgTagsUniqueViolation  = "provided feature and tags pair already exists"
	errMsgBannerNotFound       = "banner not found by provided feature id and tag (and revision_id)"
	errMsgBannerNotFoundDelete = "banner not found by provided banner_id"
	errMsgBannerConflict       = "conflict due banner update: banner with provided params already exists"

	errMsgCacheNotFound = "cache not found"
)

var (
	ErrUserNotFound          = errors.New(errMsgUserNotFound)          // ErrUserNotFound is User Not Found Error object
	ErrUserAlreadyRegistered = errors.New(errMsgUserAlreadyRegistered) // ErrUserAlreadyRegistered is User Already Registered Error object

	ErrTagsUniqueViolation  = errors.New(errMsgTagsUniqueViolation)  // ErrTagsUniqueViolation is Tags Unique Violation Error object
	ErrBannerNotFound       = errors.New(errMsgBannerNotFound)       // ErrBannerNotFound is Banner Not Found Error object
	ErrBannerNotFoundDelete = errors.New(errMsgBannerNotFoundDelete) // ErrBannerNotFoundDelete is Banner Not Found Delete Error object
	ErrBannerConflict       = errors.New(errMsgBannerConflict)       // ErrBannerConflict is Banner Conflict Error object

	ErrCacheNotFound = errors.New(errMsgCacheNotFound) // ErrCacheNotFound is Cache Not Found Error object
)
