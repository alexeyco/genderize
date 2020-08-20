package genderize

import (
	"errors"
)

var (
	// ErrResponseHeader wrong response header.
	ErrResponseHeader = errors.New("response header error")

	// ErrResponseBody wrong response body.
	ErrResponseBody = errors.New("response body error")

	// ErrInvalidAPIKey API key is invalid.
	ErrInvalidAPIKey = errors.New("invalid API key")

	// ErrSubscriptionIsNotActive subscriptions problem.
	ErrSubscriptionIsNotActive = errors.New("subscription is not active")

	// ErrValidation validation error.
	ErrValidation = errors.New("validation error")

	// ErrTooManyRequests too many requests.
	ErrTooManyRequests = errors.New("too many requests")

	// ErrInternal internal API server error.
	ErrInternal = errors.New("internal API error")

	// ErrNothingFound nothing found error.
	ErrNothingFound = errors.New("nothing found")
)
