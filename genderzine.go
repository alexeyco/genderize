package genderize

import (
	"net/http"
	"time"
)

// host genderize.io API host.
const host = "https://api.genderize.io"

// Gender type.
type Gender string

const (
	// Female gender type.
	Female Gender = "female"

	// Male gender type.
	Male Gender = "male"

	// Unknown gender type.
	Unknown Gender = "unknown"
)

// Info API info.
type Info struct {
	// Limit the amount of names available in the current time window.
	Limit int64

	// Remaining the number of names left in the current time window.
	Remaining int64

	// Reset seconds remaining until a new time window opens.
	Reset time.Duration
}

// Response API response.
type Response struct {
	Name        string  `json:"name,omitempty"`
	Gender      Gender  `json:"gender,omitempty"`
	Probability float64 `json:"probability,omitempty"`
	Count       int64   `json:"count,omitempty"`
}

// New returns new genderize.io API client instance.
func New() *Client {
	return &Client{
		httpClient: http.DefaultClient,
		info:       &Info{},
	}
}
