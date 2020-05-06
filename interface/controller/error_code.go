package controller

// ErrCode is error code.
type ErrCode string

// Server error.
const (
	InternalFailure   ErrCode = "InternalFailure"
	InternalDBFailure ErrCode = "InternalDBFailure"
)

// User error.
const (
	InvalidParameterValueFailure ErrCode = "InvalidParameterValueFailure"
	NoSuchDataFailure            ErrCode = "NoSuchDataFailure"
	AlreadyExistsFailure         ErrCode = "AlreadyExistsFailure"
)
