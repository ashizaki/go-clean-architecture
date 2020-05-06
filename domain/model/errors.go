package model

import (
	"fmt"
	"strings"
)

// AlreadyExistError expresses already specified data has existed.
type AlreadyExistError struct {
	BaseErr error
	PropertyName
	PropertyValue interface{}
	DomainModelName
}

// Error returns error message.
func (e *AlreadyExistError) Error() string {
	return fmt.Sprintf("%s, %s, is already exists", e.PropertyName, e.DomainModelName)
}

// SQLError is SQL error.
type SQLError struct {
	BaseErr       error
	InvalidReason string
}

// Error returns error message.
func (e *SQLError) Error() string {
	return e.InvalidReason
}

// InvalidParamError is inappropriate parameter error。
type InvalidParamError struct {
	BaseErr error
	PropertyName
	PropertyValue interface{}
	InvalidReason string
}

// Error returns error message.
func (e *InvalidParamError) Error() string {
	return fmt.Sprintf("%s, %v, is invalid, %s", e.PropertyName, e.PropertyValue, e.InvalidReason)
}

// InvalidParamsError is inappropriate parameters error。
type InvalidParamsError struct {
	Errors []*InvalidParamError
}

// Error returns error message.
func (e *InvalidParamsError) Error() string {
	length := len(e.Errors)
	messages := make([]string, length, length)
	for i, err := range e.Errors {
		messages[i] = err.Error()
	}
	return strings.Join(messages, ",")
}

// NoSuchDataError is not existing specified data error.
type NoSuchDataError struct {
	BaseErr error
	PropertyName
	PropertyValue interface{}
	DomainModelName
}

// Error returns error message.
func (e *NoSuchDataError) Error() string {
	return fmt.Sprintf("no such data, %s: %v, %s", e.PropertyName, e.PropertyValue, e.DomainModelName)
}

// RepositoryError is Repository error.
type RepositoryError struct {
	BaseErr          error
	RepositoryMethod RepositoryMethod
	DomainModelName
}

// Error returns error message.
func (e *RepositoryError) Error() string {
	return fmt.Sprintf("failed Repository operation, %s, %s", e.RepositoryMethod, e.DomainModelName)
}
