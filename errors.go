package genderzine

import "errors"

var (
	// ErrInvalidAPIKey invalid API key error.
	ErrInvalidAPIKey = errors.New("invalid API key")

	// ErrSubscriptionIsNotActive subscription is not active error.
	ErrSubscriptionIsNotActive = errors.New("subscription is not active")

	// ErrMissingName missing 'name' parameter error.
	ErrMissingName = errors.New("missing 'name' parameter")

	// ErrInvalidName invalid 'name' parameter error.
	ErrInvalidName = errors.New("invalid 'name' parameter")

	// ErrRequestLimitReached request limit reached error.
	ErrRequestLimitReached = errors.New("request limit reached")

	// ErrRequestLimitTooLow request limit too low to process request error.
	ErrRequestLimitTooLow = errors.New("request limit too low to process request")

	// ErrUnknown unknown error.
	ErrUnknown = errors.New("something went wrong")
)
