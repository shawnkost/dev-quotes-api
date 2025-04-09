package errors

import (
	"net/http"
)

// ErrorType represents the type of error
type ErrorType string

const (
	// ErrorTypeNotFound represents a resource not found error
	ErrorTypeNotFound ErrorType = "NOT_FOUND"
	// ErrorTypeInternal represents an internal server error
	ErrorTypeInternal ErrorType = "INTERNAL"
	// ErrorTypeValidation represents a validation error
	ErrorTypeValidation ErrorType = "VALIDATION"
)

// APIError represents a standardized API error response
type APIError struct {
	Type    ErrorType `json:"type"`
	Message string    `json:"message"`
	Code    int       `json:"code"`
}

// Error implements the error interface
func (e *APIError) Error() string {
	return e.Message
}

// NewNotFoundError creates a new not found error
func NewNotFoundError(message string) *APIError {
	return &APIError{
		Type:    ErrorTypeNotFound,
		Message: message,
		Code:    http.StatusNotFound,
	}
}

// NewInternalError creates a new internal server error
func NewInternalError(message string) *APIError {
	return &APIError{
		Type:    ErrorTypeInternal,
		Message: message,
		Code:    http.StatusInternalServerError,
	}
}

// NewValidationError creates a new validation error
func NewValidationError(message string) *APIError {
	return &APIError{
		Type:    ErrorTypeValidation,
		Message: message,
		Code:    http.StatusBadRequest,
	}
}

// IsNotFound checks if an error is a not found error
func IsNotFound(err error) bool {
	if apiErr, ok := err.(*APIError); ok {
		return apiErr.Type == ErrorTypeNotFound
	}
	return false
}

// IsInternal checks if an error is an internal error
func IsInternal(err error) bool {
	if apiErr, ok := err.(*APIError); ok {
		return apiErr.Type == ErrorTypeInternal
	}
	return false
}

// IsValidation checks if an error is a validation error
func IsValidation(err error) bool {
	if apiErr, ok := err.(*APIError); ok {
		return apiErr.Type == ErrorTypeValidation
	}
	return false
}
