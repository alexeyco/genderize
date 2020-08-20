package genderize

// Response API response.
type Response struct {
	Name        string  `json:"name,omitempty"`
	Gender      string  `json:"gender,omitempty"`
	Probability float64 `json:"probability,omitempty"`
	Count       int64   `json:"count,omitempty"`
}

// Error response error.
type Error struct {
	Error string `json:"error,omitempty"`
}

type apiResponse struct {
	Response
	Error
}
