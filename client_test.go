package genderize_test

import (
	"context"
	"errors"
	"io/ioutil"
	"net/http"
	"strings"
	"testing"

	"github.com/alexeyco/genderize"
)

type testClientHandler func(req *http.Request) (res *http.Response, err error)

type testClientRoundTripper struct {
	fn testClientHandler
}

func (t testClientRoundTripper) RoundTrip(req *http.Request) (*http.Response, error) {
	return t.fn(req)
}

func testClientClient(fn testClientHandler) *http.Client {
	return &http.Client{
		Transport: &testClientRoundTripper{
			fn: fn,
		},
	}
}

// nolint:gochecknoglobals,golint,stylecheck
var testClientErr = errors.New("wtf")

func TestClient_Execute_Err(t *testing.T) {
	httpClient := testClientClient(func(_ *http.Request) (res *http.Response, err error) {
		return nil, testClientErr
	})

	r := genderize.NewRequest(context.TODO()).
		Name("Alice", "John")

	_, err := genderize.NewClient(genderize.WithHTTPClient(httpClient)).
		Execute(r)
	if err == nil {
		t.Error(`Should not be nil`)
	}
}

// nolint:dupl
func TestClient_Execute_ErrInvalidAPIKey(t *testing.T) {
	httpClient := testClientClient(func(_ *http.Request) (res *http.Response, err error) {
		h := http.Header{}
		h.Set(genderize.HdrXRateLimitLimit, "0")
		h.Set(genderize.HdrXRateLimitRemaining, "0")
		h.Set(genderize.HdrXRateReset, "0")

		res = &http.Response{
			StatusCode: http.StatusUnauthorized,
			Header:     h,
			Body:       ioutil.NopCloser(strings.NewReader(`{}`)),
		}

		return
	})

	r := genderize.NewRequest(context.TODO()).
		Name("Alice", "John")

	_, err := genderize.NewClient(genderize.WithHTTPClient(httpClient)).
		Execute(r)
	if err == nil {
		t.Error(`Should not be nil`)
	}

	if !errors.Is(err, genderize.ErrInvalidAPIKey) {
		t.Error(`Should be genderize.ErrInvalidAPIKey`)
	}
}

// nolint:dupl
func TestClient_Execute_ErrSubscriptionIsNotActive(t *testing.T) {
	httpClient := testClientClient(func(_ *http.Request) (res *http.Response, err error) {
		h := http.Header{}
		h.Set(genderize.HdrXRateLimitLimit, "0")
		h.Set(genderize.HdrXRateLimitRemaining, "0")
		h.Set(genderize.HdrXRateReset, "0")

		res = &http.Response{
			StatusCode: http.StatusPaymentRequired,
			Header:     h,
			Body:       ioutil.NopCloser(strings.NewReader(`{}`)),
		}

		return
	})

	r := genderize.NewRequest(context.TODO()).
		Name("Alice", "John")

	_, err := genderize.NewClient(genderize.WithHTTPClient(httpClient)).
		Execute(r)
	if err == nil {
		t.Error(`Should not be nil`)
	}

	if !errors.Is(err, genderize.ErrSubscriptionIsNotActive) {
		t.Error(`Should be genderize.ErrSubscriptionIsNotActive`)
	}
}

// nolint:dupl
func TestClient_Execute_ErrValidation(t *testing.T) {
	httpClient := testClientClient(func(_ *http.Request) (res *http.Response, err error) {
		h := http.Header{}
		h.Set(genderize.HdrXRateLimitLimit, "0")
		h.Set(genderize.HdrXRateLimitRemaining, "0")
		h.Set(genderize.HdrXRateReset, "0")

		res = &http.Response{
			StatusCode: http.StatusUnprocessableEntity,
			Header:     h,
			Body:       ioutil.NopCloser(strings.NewReader(`{}`)),
		}

		return
	})

	r := genderize.NewRequest(context.TODO()).
		Name("Alice", "John")

	_, err := genderize.NewClient(genderize.WithHTTPClient(httpClient)).
		Execute(r)
	if err == nil {
		t.Error(`Should not be nil`)
	}

	if !errors.Is(err, genderize.ErrValidation) {
		t.Error(`Should be genderize.ErrValidation`)
	}
}

// nolint:dupl
func TestClient_Execute_ErrTooManyRequests(t *testing.T) {
	httpClient := testClientClient(func(_ *http.Request) (res *http.Response, err error) {
		h := http.Header{}
		h.Set(genderize.HdrXRateLimitLimit, "0")
		h.Set(genderize.HdrXRateLimitRemaining, "0")
		h.Set(genderize.HdrXRateReset, "0")

		res = &http.Response{
			StatusCode: http.StatusTooManyRequests,
			Header:     h,
			Body:       ioutil.NopCloser(strings.NewReader(`{}`)),
		}

		return
	})

	r := genderize.NewRequest(context.TODO()).
		Name("Alice", "John")

	_, err := genderize.NewClient(genderize.WithHTTPClient(httpClient)).
		Execute(r)
	if err == nil {
		t.Error(`Should not be nil`)
	}

	if !errors.Is(err, genderize.ErrTooManyRequests) {
		t.Error(`Should be genderize.ErrTooManyRequests`)
	}
}

func TestClient_ExecuteX(t *testing.T) {
	httpClient := testClientClient(func(_ *http.Request) (res *http.Response, err error) {
		h := http.Header{}
		h.Set(genderize.HdrXRateLimitLimit, "0")
		h.Set(genderize.HdrXRateLimitRemaining, "0")
		h.Set(genderize.HdrXRateReset, "0")

		res = &http.Response{
			StatusCode: http.StatusOK,
			Header:     h,
			Body:       ioutil.NopCloser(strings.NewReader(`[]`)),
		}

		return
	})

	r := genderize.NewRequest(context.TODO()).
		Name("Alice", "John")

	defer func() {
		if err := recover(); err != nil {
			t.Errorf(`Should be nil, "%s" given`, err)
		}
	}()

	_ = genderize.NewClient(genderize.WithHTTPClient(httpClient)).
		ExecuteX(r)
}

func TestClient_ExecuteX_Panic(t *testing.T) {
	httpClient := testClientClient(func(_ *http.Request) (res *http.Response, err error) {
		return nil, testClientErr
	})

	r := genderize.NewRequest(context.TODO()).
		Name("Alice", "John")

	defer func() {
		if err := recover(); err == nil {
			t.Error(`Should not be nil`)
		}
	}()

	_ = genderize.NewClient(genderize.WithHTTPClient(httpClient)).
		ExecuteX(r)
}
