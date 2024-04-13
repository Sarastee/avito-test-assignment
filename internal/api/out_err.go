package api

import "errors"

const (
	errMsgInternalError = "something went wrong, we are already working on it"

	errMsgTagNotFound     = "tag_id not found"
	errMsgTagIsNotANumber = "tag_id is not a number"

	errMsgFeatureNotFound     = "feature_id not found"
	errMsgFeatureIsNotANumber = "feature_id is not a number"

	errMsgRevisionIsNotANumber = "revision_id is not a number"

	errMsgLimitIsNotANumber  = "limit is not a number"
	errMsgOffsetIsNotANumber = "offset is not a number"

	errMsgInsufficientRights = "insufficient rights to execute the command"

	errMsgNoTokenProvided = "no token provided"
	errMsgInvalidToken    = "invalid token"
)

var (
	ErrInternalError = errors.New(errMsgInternalError) // ErrInternalError is Internal Server Error object

	ErrTagNotFound     = errors.New(errMsgTagNotFound)     // ErrTagNotFound is Tag Not Found Error object
	ErrTagIsNotANumber = errors.New(errMsgTagIsNotANumber) // ErrTagIsNotANumber is Tag Is Not A Number Error object

	ErrFeatureNotFound     = errors.New(errMsgFeatureNotFound)     // ErrFeatureNotFound is Feature Not Found Error object
	ErrFeatureIsNotANumber = errors.New(errMsgFeatureIsNotANumber) // ErrFeatureIsNotANumber is Feature Is Not A Number Error

	ErrRevisionIsNotANumber = errors.New(errMsgRevisionIsNotANumber) // ErrRevisionIsNotANumber is Revision Is Not A Number Error object

	ErrLimitIsNotANumber  = errors.New(errMsgLimitIsNotANumber)  // ErrLimitIsNotANumber is Limit Is Not A Number Error object
	ErrOffsetIsNotANumber = errors.New(errMsgOffsetIsNotANumber) // ErrOffsetIsNotANumber is Offset Is Not A Number Error object
	ErrInsufficientRights = errors.New(errMsgInsufficientRights) // ErrInsufficientRights is Insufficient Rights Error object

	ErrNoTokenProvided = errors.New(errMsgNoTokenProvided) // ErrNoTokenProvided is No Token Provided Error object
	ErrInvalidToken    = errors.New(errMsgInvalidToken)    // ErrInvalidToken is Invalid Token Error object
)
