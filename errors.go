package genderize

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

	// ErrEmptyXRateLimitLimit empty X-Rate-Limit-Limit response header.
	ErrEmptyXRateLimitLimit = errors.New("empty X-Rate-Limit-Limit response header")

	// ErrEmptyXRateLimitRemaining empty X-Rate-Limit-Limit response header.
	ErrEmptyXRateLimitRemaining = errors.New("empty X-Rate-Limit-Remaining response header")

	// ErrEmptyXRateReset empty X-Rate-Reset response header.
	ErrEmptyXRateReset = errors.New("empty X-Rate-Reset response header")

	// ErrWrongXRateLimitLimit wrong X-Rate-Limit-Limit response header.
	ErrWrongXRateLimitLimit = errors.New("wrong X-Rate-Limit-Limit response header")

	// ErrWrongXRateLimitRemaining wrong X-Rate-Limit-Remaining response header.
	ErrWrongXRateLimitRemaining = errors.New("wrong X-Rate-Limit-Remaining response header")

	// ErrWrongXRateReset wrong X-Rate-Reset response header.
	ErrWrongXRateReset = errors.New("wrong X-Rate-Reset response header")

	// ErrUnknown unknown error.
	ErrUnknown = errors.New("something went wrong")
)
