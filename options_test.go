package genderize_test

import (
	"net/http"
	"testing"

	"github.com/alexeyco/genderize"
)

func TestWithAPIKey(t *testing.T) {
	o := &genderize.Options{}

	genderize.WithAPIKey("FooBar")(o)

	if o.APIKey != "FooBar" {
		t.Errorf(`Should be "%s", "%s" given`, "FooBar", o.APIKey)
	}
}

func TestWithHTTPClient(t *testing.T) {
	o := &genderize.Options{}

	genderize.WithHTTPClient(http.DefaultClient)(o)

	if o.HTTPClient == nil {
		t.Error(`Should not be nil`)
	}
}
