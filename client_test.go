package genderize_test

import (
	"context"
	"errors"
	"net/http"
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

func TestClient_Execute_Err(t *testing.T) {
	httpClient := testClientClient(func(_ *http.Request) (res *http.Response, err error) {
		err = errors.New("wtf")

		return
	})

	r := genderize.NewRequest(context.TODO()).
		Name("Alice", "John")

	_, err := genderize.NewClient(genderize.WithHTTPClient(httpClient)).
		Execute(r)
	if err == nil {
		t.Error(`Should not be nil`)
	}
}

func TestClient_ExecuteX(t *testing.T) {
	httpClient := testClientClient(func(_ *http.Request) (res *http.Response, err error) {
		err = errors.New("wtf")

		return
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
