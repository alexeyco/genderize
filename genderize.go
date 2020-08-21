package genderize

const (
	// HdrXRateLimitLimit the amount of names available in the current time window.
	HdrXRateLimitLimit = "X-Rate-Limit-Limit"

	// HdrXRateLimitRemaining the number of names left in the current time window.
	HdrXRateLimitRemaining = "X-Rate-Limit-Remaining"

	// HdrXRateReset seconds remaining until a new time window opens.
	HdrXRateReset = "X-Rate-Reset"
)
