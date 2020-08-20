package genderize

import (
	"net/url"
	"sync"
)

const endpoint = "https://api.genderize.io"

// Request API request.
type Request struct {
	mu    sync.Mutex
	url   *url.URL
	query url.Values
}

// Name sets person names.
func (r *Request) Name(name ...string) *Request {
	r.mu.Lock()
	defer r.mu.Unlock()

	for _, n := range name {
		r.query.Add("name", n)
	}

	return r
}

// CountryID sets country ISO 3166-1 alpha-2 ID.
func (r *Request) CountryID(countryID string) *Request {
	r.mu.Lock()
	defer r.mu.Unlock()

	r.query.Set("country_id", countryID)

	return r
}

// Encode returns request URL.
func (r *Request) Encode(apiKey ...string) string {
	r.mu.Lock()
	defer r.mu.Unlock()

	if len(apiKey) != 0 {
		r.query.Set("apikey", apiKey[0])
	}

	r.url.RawQuery = r.query.Encode()

	return r.url.String()
}

// NewRequest returns new request instance.
func NewRequest() *Request {
	r := &Request{
		query: url.Values{},
	}

	r.url, _ = url.Parse(endpoint)

	return r
}
