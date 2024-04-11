package repository

import "errors"

const (
	errMsgNoRowsAffected      = "no rows were affected"
	errMsgTagsUniqueViolation = "provided feature and tags pair already exists"
)

var (
	ErrNoRowsAffected      = errors.New(errMsgNoRowsAffected)      // ErrNoRowsAffected is No Rows Affected Error object
	ErrTagsUniqueViolation = errors.New(errMsgTagsUniqueViolation) // ErrTagsUniqueViolation is Tags Unique Violation Error object
)
