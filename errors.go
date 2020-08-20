package genderize

import (
	"errors"
)

var (
	// ErrResponseHeader wrong response header.
	ErrResponseHeader = errors.New("response header error")

	// ErrResponseBody wrong response body.
	ErrResponseBody = errors.New("response body error")

	// ErrInvalidAPIKey
	ErrInvalidAPIKey = errors.New("invalid API key")

	// ErrSubscriptionIsNotActive
	ErrSubscriptionIsNotActive = errors.New("subscription is not active")

	// ErrValidation
	ErrValidation = errors.New("validation error")

	// ErrTooManyRequests
	ErrTooManyRequests = errors.New("too many requests")

	// ErrInternal
	ErrInternal = errors.New("internal API error")

	// ErrNothingFound nothing found error.
	ErrNothingFound = errors.New("nothing found")
)
