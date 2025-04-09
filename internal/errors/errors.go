package errors

import (
	"net/http"
)

type ErrorType string

const (
	ErrorTypeNotFound   ErrorType = "NOT_FOUND"
	ErrorTypeInternal   ErrorType = "INTERNAL"
	ErrorTypeValidation ErrorType = "VALIDATION"
)

type APIError struct {
	Type    ErrorType `json:"type"`
	Message string    `json:"message"`
	Code    int       `json:"code"`
}

func (e *APIError) Error() string {
	return e.Message
}

func NewNotFoundError(message string) *APIError {
	return &APIError{
		Type:    ErrorTypeNotFound,
		Message: message,
		Code:    http.StatusNotFound,
	}
}

func NewInternalError(message string) *APIError {
	return &APIError{
		Type:    ErrorTypeInternal,
		Message: message,
		Code:    http.StatusInternalServerError,
	}
}

func NewValidationError(message string) *APIError {
	return &APIError{
		Type:    ErrorTypeValidation,
		Message: message,
		Code:    http.StatusBadRequest,
	}
}

func IsNotFound(err error) bool {
	if apiErr, ok := err.(*APIError); ok {
		return apiErr.Type == ErrorTypeNotFound
	}
	return false
}

func IsInternal(err error) bool {
	if apiErr, ok := err.(*APIError); ok {
		return apiErr.Type == ErrorTypeInternal
	}
	return false
}

func IsValidation(err error) bool {
	if apiErr, ok := err.(*APIError); ok {
		return apiErr.Type == ErrorTypeValidation
	}
	return false
}
