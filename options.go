package genderize

import "net/http"

// Options of a client.
type Options struct {
	APIKey     string
	HTTPClient *http.Client
}

// Option callback.
type Option func(o *Options)

// WithAPIKey sets API key.
func WithAPIKey(apiKey string) Option {
	return func(o *Options) {
		o.APIKey = apiKey
	}
}

// WithHTTPClient sets custom HTTP client.
func WithHTTPClient(httpClient *http.Client) Option {
	return func(o *Options) {
		o.HTTPClient = httpClient
	}
}
