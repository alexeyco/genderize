package genderize

import "net/http"

// Options of a client.
type Options struct {
	apiKey     string
	httpClient *http.Client
}

// Option callback.
type Option func(o *Options)

// WithAPIKey sets API key.
func WithAPIKey(apiKey string) Option {
	return func(o *Options) {
		o.apiKey = apiKey
	}
}

// WithHTTPClient sets custom HTTP client.
func WithHTTPClient(httpClient *http.Client) Option {
	return func(o *Options) {
		o.httpClient = httpClient
	}
}
