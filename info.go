package genderize

import "time"

// Info API rate limits info.
type Info struct {
	Limit     int64
	Remaining int64
	Reset     time.Duration
}
