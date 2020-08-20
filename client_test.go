package genderize_test

import (
	"context"
	"errors"
	"io/ioutil"
	"net/http"
	"strings"
	"testing"
	"time"

	"github.com/alexeyco/genderize"
)

type roundTripper struct {
	fn func(*http.Request) (*http.Response, error)
}

func (r *roundTripper) RoundTrip(req *http.Request) (res *http.Response, err error) {
	return r.fn(req)
}

func newClient(fn func(*http.Request) (*http.Response, error)) *http.Client {
	return &http.Client{
		Transport: &roundTripper{
			fn: fn,
		},
	}
}

func TestClient_SetAPIKey_Empty(t *testing.T) {
	_, err := genderize.New().
		SetHTTPClient(newClient(func(req *http.Request) (res *http.Response, err error) {
			apiKey := req.URL.Query().Get("apikey")

			if apiKey != "" {
				t.Errorf(`API key should be empty, "%s" given`, apiKey)
			}

			header := http.Header{}
			header.Add("X-Rate-Limit-Limit", "0")
			header.Add("X-Rate-Limit-Remaining", "0")
			header.Add("X-Rate-Reset", "0")

			body := strings.NewReader(`{"name":"Name","gender":"male","probability":0,"count":0}`)

			res = &http.Response{
				StatusCode: http.StatusOK,
				Body:       ioutil.NopCloser(body),
				Header:     header,
			}

			return
		})).
		Check(context.Background(), "Alice")
	if err != nil {
		t.Errorf(`Error should be nil, "%s" given`, err)
	}
}

func TestClient_SetAPIKey_NonEmpty(t *testing.T) {
	_, err := genderize.New().
		SetAPIKey("FooBar").
		SetHTTPClient(newClient(func(req *http.Request) (res *http.Response, err error) {
			apiKey := req.URL.Query().Get("apikey")

			if apiKey != "FooBar" {
				t.Errorf(`API key should be "%s", "%s" given`, "FooBar", apiKey)
			}

			header := http.Header{}
			header.Add("X-Rate-Limit-Limit", "0")
			header.Add("X-Rate-Limit-Remaining", "0")
			header.Add("X-Rate-Reset", "0")

			body := strings.NewReader(`{"name":"Name","gender":"male","probability":0,"count":0}`)

			res = &http.Response{
				StatusCode: http.StatusOK,
				Body:       ioutil.NopCloser(body),
				Header:     header,
			}

			return
		})).
		Check(context.Background(), "Alice")
	if err != nil {
		t.Errorf(`Error should be nil, "%s" given`, err)
	}
}

func TestClient_Check_ErrInvalidAPIKey(t *testing.T) {
	client := genderize.New()

	_, err := client.SetHTTPClient(newClient(func(req *http.Request) (res *http.Response, err error) {
		header := http.Header{}
		header.Add("X-Rate-Limit-Limit", "123")
		header.Add("X-Rate-Limit-Remaining", "456")
		header.Add("X-Rate-Reset", "789")

		body := strings.NewReader(`{"error":"Invalid API key"}`)

		res = &http.Response{
			StatusCode: http.StatusUnauthorized,
			Body:       ioutil.NopCloser(body),
			Header:     header,
		}

		return
	})).Check(context.Background(), "Alice")
	if err == nil {
		t.Error(`Error should not be nil`)
	}

	if !errors.Is(err, genderize.ErrInvalidAPIKey) {
		t.Errorf(`Error should be "%s", "%s" given`, "genderize.ErrInvalidAPIKey", err)
	}
}

func TestClient_Check_ErrSubscriptionIsNotActive(t *testing.T) {
	client := genderize.New()

	_, err := client.SetHTTPClient(newClient(func(req *http.Request) (res *http.Response, err error) {
		header := http.Header{}
		header.Add("X-Rate-Limit-Limit", "123")
		header.Add("X-Rate-Limit-Remaining", "456")
		header.Add("X-Rate-Reset", "789")

		body := strings.NewReader(`{"error":"Subscription is not active"}`)

		res = &http.Response{
			StatusCode: http.StatusPaymentRequired,
			Body:       ioutil.NopCloser(body),
			Header:     header,
		}

		return
	})).Check(context.Background(), "Alice")
	if err == nil {
		t.Error(`Error should not be nil`)
	}

	if !errors.Is(err, genderize.ErrSubscriptionIsNotActive) {
		t.Errorf(`Error should be "%s", "%s" given`, "genderize.ErrSubscriptionIsNotActive", err)
	}
}

func TestClient_Check_ErrMissingName(t *testing.T) {
	client := genderize.New()

	_, err := client.SetHTTPClient(newClient(func(req *http.Request) (res *http.Response, err error) {
		header := http.Header{}
		header.Add("X-Rate-Limit-Limit", "123")
		header.Add("X-Rate-Limit-Remaining", "456")
		header.Add("X-Rate-Reset", "789")

		body := strings.NewReader(`{"error":"Missing 'name' parameter"}`)

		res = &http.Response{
			StatusCode: http.StatusUnprocessableEntity,
			Body:       ioutil.NopCloser(body),
			Header:     header,
		}

		return
	})).Check(context.Background(), "Alice")
	if err == nil {
		t.Error(`Error should not be nil`)
	}

	if !errors.Is(err, genderize.ErrMissingName) {
		t.Errorf(`Error should be "%s", "%s" given`, "genderize.ErrMissingName", err)
	}
}

func TestClient_Check_ErrInvalidName(t *testing.T) {
	client := genderize.New()

	_, err := client.SetHTTPClient(newClient(func(req *http.Request) (res *http.Response, err error) {
		header := http.Header{}
		header.Add("X-Rate-Limit-Limit", "123")
		header.Add("X-Rate-Limit-Remaining", "456")
		header.Add("X-Rate-Reset", "789")

		body := strings.NewReader(`{"error":"Invalid 'name' parameter"}`)

		res = &http.Response{
			StatusCode: http.StatusUnprocessableEntity,
			Body:       ioutil.NopCloser(body),
			Header:     header,
		}

		return
	})).Check(context.Background(), "Alice")
	if err == nil {
		t.Error(`Error should not be nil`)
	}

	if !errors.Is(err, genderize.ErrInvalidName) {
		t.Errorf(`Error should be "%s", "%s" given`, "genderize.ErrInvalidName", err)
	}
}

func TestClient_Check_ErrRequestLimitReached(t *testing.T) {
	client := genderize.New()

	_, err := client.SetHTTPClient(newClient(func(req *http.Request) (res *http.Response, err error) {
		header := http.Header{}
		header.Add("X-Rate-Limit-Limit", "123")
		header.Add("X-Rate-Limit-Remaining", "456")
		header.Add("X-Rate-Reset", "789")

		body := strings.NewReader(`{"error":"Request limit reached"}`)

		res = &http.Response{
			StatusCode: http.StatusTooManyRequests,
			Body:       ioutil.NopCloser(body),
			Header:     header,
		}

		return
	})).Check(context.Background(), "Alice")
	if err == nil {
		t.Error(`Error should not be nil`)
	}

	if !errors.Is(err, genderize.ErrRequestLimitReached) {
		t.Errorf(`Error should be "%s", "%s" given`, "genderize.ErrRequestLimitReached", err)
	}
}

func TestClient_Check_ErrRequestLimitTooLow(t *testing.T) {
	client := genderize.New()

	_, err := client.SetHTTPClient(newClient(func(req *http.Request) (res *http.Response, err error) {
		header := http.Header{}
		header.Add("X-Rate-Limit-Limit", "123")
		header.Add("X-Rate-Limit-Remaining", "456")
		header.Add("X-Rate-Reset", "789")

		body := strings.NewReader(`{"error":"Request limit too low to process request"}`)

		res = &http.Response{
			StatusCode: http.StatusTooManyRequests,
			Body:       ioutil.NopCloser(body),
			Header:     header,
		}

		return
	})).Check(context.Background(), "Alice")
	if err == nil {
		t.Error(`Error should not be nil`)
	}

	if !errors.Is(err, genderize.ErrRequestLimitTooLow) {
		t.Errorf(`Error should be "%s", "%s" given`, "genderize.ErrRequestLimitTooLow", err)
	}
}

func TestClient_Check_Ok(t *testing.T) {
	client := genderize.New()

	res, err := client.SetHTTPClient(newClient(func(req *http.Request) (res *http.Response, err error) {
		header := http.Header{}
		header.Add("X-Rate-Limit-Limit", "123")
		header.Add("X-Rate-Limit-Remaining", "456")
		header.Add("X-Rate-Reset", "789")

		body := strings.NewReader(`{"name":"Name","gender":"male","probability":0.1,"count":5}`)

		res = &http.Response{
			StatusCode: http.StatusOK,
			Body:       ioutil.NopCloser(body),
			Header:     header,
		}

		return
	})).Check(context.Background(), "Alice")
	if err != nil {
		t.Errorf(`Error should be nil, "%s" given`, err)
	}

	if res.Name != "Name" {
		t.Errorf(`Name should be "%s", "%s" given`, "Name", res.Name)
	}

	if res.Gender != "male" {
		t.Errorf(`Gender should be "%s", "%s" given`, "male", res.Gender)
	}

	if res.Probability != .1 {
		t.Errorf(`Gender should be %f, %f given`, .1, res.Probability)
	}

	if res.Count != 5 {
		t.Errorf(`Count should be %d, %d given`, 5, res.Count)
	}
}

func TestClient_Info_EmptyXRateLimitLimit(t *testing.T) {
	client := genderize.New()

	_, err := client.SetHTTPClient(newClient(func(req *http.Request) (res *http.Response, err error) {
		header := http.Header{}
		header.Add("X-Rate-Limit-Remaining", "456")
		header.Add("X-Rate-Reset", "789")

		body := strings.NewReader(`{"name":"Name","gender":"male","probability":0,"count":0}`)

		res = &http.Response{
			StatusCode: http.StatusOK,
			Body:       ioutil.NopCloser(body),
			Header:     header,
		}

		return
	})).Check(context.Background(), "Alice")

	if err == nil {
		t.Error(`Error should not be nil, nil given`)
	}

	if !errors.Is(err, genderize.ErrEmptyXRateLimitLimit) {
		t.Errorf(`Error should be "%s"`, "genderize.ErrEmptyXRateLimitLimit")
	}
}

func TestClient_Info_WrongXRateLimitLimit(t *testing.T) {
	client := genderize.New()

	_, err := client.SetHTTPClient(newClient(func(req *http.Request) (res *http.Response, err error) {
		header := http.Header{}
		header.Add("X-Rate-Limit-Limit", "abc")
		header.Add("X-Rate-Limit-Remaining", "456")
		header.Add("X-Rate-Reset", "789")

		body := strings.NewReader(`{"name":"Name","gender":"male","probability":0,"count":0}`)

		res = &http.Response{
			StatusCode: http.StatusOK,
			Body:       ioutil.NopCloser(body),
			Header:     header,
		}

		return
	})).Check(context.Background(), "Alice")

	if err == nil {
		t.Error(`Error should not be nil, nil given`)
	}

	if !errors.Is(err, genderize.ErrWrongXRateLimitLimit) {
		t.Errorf(`Error should be "%s"`, "genderize.ErrWrongXRateLimitLimit")
	}
}

func TestClient_Info_EmptyXRateLimitRemaining(t *testing.T) {
	client := genderize.New()

	_, err := client.SetHTTPClient(newClient(func(req *http.Request) (res *http.Response, err error) {
		header := http.Header{}
		header.Add("X-Rate-Limit-Limit", "123")
		header.Add("X-Rate-Reset", "789")

		body := strings.NewReader(`{"name":"Name","gender":"male","probability":0,"count":0}`)

		res = &http.Response{
			StatusCode: http.StatusOK,
			Body:       ioutil.NopCloser(body),
			Header:     header,
		}

		return
	})).Check(context.Background(), "Alice")

	if err == nil {
		t.Error(`Error should not be nil, nil given`)
	}

	if !errors.Is(err, genderize.ErrEmptyXRateLimitRemaining) {
		t.Errorf(`Error should be "%s"`, "genderize.ErrEmptyXRateLimitRemaining")
	}
}

func TestClient_Info_WrongXRateLimitRemaining(t *testing.T) {
	client := genderize.New()

	_, err := client.SetHTTPClient(newClient(func(req *http.Request) (res *http.Response, err error) {
		header := http.Header{}
		header.Add("X-Rate-Limit-Limit", "123")
		header.Add("X-Rate-Limit-Remaining", "abc")
		header.Add("X-Rate-Reset", "789")

		body := strings.NewReader(`{"name":"Name","gender":"male","probability":0,"count":0}`)

		res = &http.Response{
			StatusCode: http.StatusOK,
			Body:       ioutil.NopCloser(body),
			Header:     header,
		}

		return
	})).Check(context.Background(), "Alice")

	if err == nil {
		t.Error(`Error should not be nil, nil given`)
	}

	if !errors.Is(err, genderize.ErrWrongXRateLimitRemaining) {
		t.Errorf(`Error should be "%s"`, "genderize.ErrWrongXRateLimitRemaining")
	}
}

func TestClient_Info_EmptyXRateReset(t *testing.T) {
	client := genderize.New()

	_, err := client.SetHTTPClient(newClient(func(req *http.Request) (res *http.Response, err error) {
		header := http.Header{}
		header.Add("X-Rate-Limit-Limit", "123")
		header.Add("X-Rate-Limit-Remaining", "456")

		body := strings.NewReader(`{"name":"Name","gender":"male","probability":0,"count":0}`)

		res = &http.Response{
			StatusCode: http.StatusOK,
			Body:       ioutil.NopCloser(body),
			Header:     header,
		}

		return
	})).Check(context.Background(), "Alice")

	if err == nil {
		t.Error(`Error should not be nil, nil given`)
	}

	if !errors.Is(err, genderize.ErrEmptyXRateReset) {
		t.Errorf(`Error should be "%s"`, "genderize.ErrEmptyXRateReset")
	}
}

func TestClient_Info_WrongXRateReset(t *testing.T) {
	client := genderize.New()

	_, err := client.SetHTTPClient(newClient(func(req *http.Request) (res *http.Response, err error) {
		header := http.Header{}
		header.Add("X-Rate-Limit-Limit", "123")
		header.Add("X-Rate-Limit-Remaining", "456")
		header.Add("X-Rate-Reset", "abc")

		body := strings.NewReader(`{"name":"Name","gender":"male","probability":0,"count":0}`)

		res = &http.Response{
			StatusCode: http.StatusOK,
			Body:       ioutil.NopCloser(body),
			Header:     header,
		}

		return
	})).Check(context.Background(), "Alice")

	if err == nil {
		t.Error(`Error should not be nil, nil given`)
	}

	if !errors.Is(err, genderize.ErrWrongXRateReset) {
		t.Errorf(`Error should be "%s"`, "genderize.ErrWrongXRateReset")
	}
}

func TestClient_Info_Ok(t *testing.T) {
	client := genderize.New()

	_, err := client.SetHTTPClient(newClient(func(req *http.Request) (res *http.Response, err error) {
		header := http.Header{}
		header.Add("X-Rate-Limit-Limit", "123")
		header.Add("X-Rate-Limit-Remaining", "456")
		header.Add("X-Rate-Reset", "789")

		body := strings.NewReader(`{"name":"Name","gender":"male","probability":0,"count":0}`)

		res = &http.Response{
			StatusCode: http.StatusOK,
			Body:       ioutil.NopCloser(body),
			Header:     header,
		}

		return
	})).Check(context.Background(), "Alice")
	if err != nil {
		t.Errorf(`Error should be nil, "%s" given`, err)
	}

	info := client.Info()

	if info.Limit != 123 {
		t.Errorf(`Should be %d, %d given`, 123, info.Limit)
	}

	if info.Remaining != 456 {
		t.Errorf(`Should be %d, %d given`, 456, info.Remaining)
	}

	if info.Reset != 789*time.Second {
		t.Errorf(`Should be %s, %s given`, 789*time.Second, info.Reset)
	}
}
