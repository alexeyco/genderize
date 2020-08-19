package genderzine

import (
	"encoding/json"
	"net/http"
	"net/url"
	"strconv"
	"sync"
	"time"

	"github.com/pkg/errors"
)

type apiResponse struct {
	Response
	Error string `json:"error,omitempty"`
}

// Client of genderzine.io.
type Client struct {
	muHTTPClient sync.Mutex
	httpClient   *http.Client

	muAPIKey sync.Mutex
	apiKey   string

	muInfo sync.Mutex
	info   *Info
}

// SetHTTPClient HTTP client setter.
func (c *Client) SetHTTPClient(httpClient *http.Client) *Client {
	c.muHTTPClient.Lock()
	defer c.muHTTPClient.Unlock()

	c.httpClient = httpClient

	return c
}

// SetAPIKey API key setter.
func (c *Client) SetAPIKey(apiKey string) *Client {
	c.muAPIKey.Lock()
	defer c.muAPIKey.Unlock()

	c.apiKey = apiKey

	return c
}

// Info returns API info.
func (c *Client) Info() *Info {
	c.muInfo.Lock()
	defer c.muInfo.Unlock()

	return c.info
}

func (c *Client) url(name string) (s string, err error) {
	c.muAPIKey.Lock()
	defer c.muAPIKey.Unlock()

	u, err := url.Parse(host)
	if err != nil {
		return
	}

	q := url.Values{}
	q.Add("name", name)

	if c.apiKey != "" {
		q.Add("apikey", c.apiKey)
	}

	u.RawQuery = q.Encode()
	s = u.String()

	return
}

func (c *Client) request(u string) (r *Response, err error) {
	c.muHTTPClient.Lock()
	defer c.muHTTPClient.Unlock()

	response, err := c.httpClient.Get(u)
	if err != nil {
		return nil, errors.Wrapf(err, "can't get HTTP request to %s", u)
	}

	defer func() {
		_ = response.Body.Close()
	}()

	var res apiResponse
	if err = json.NewDecoder(response.Body).Decode(&res); err != nil {
		return nil, errors.Wrap(err, "can't decode response body")
	}

	c.muInfo.Lock()
	defer c.muInfo.Unlock()

	h := response.Header
	if c.info.Limit, err = strconv.ParseInt(h.Get("X-Rate-Limit-Limit"), 10, 64); err != nil {
		return nil, errors.Wrapf(err, `can't parse "X-Rate-Limit-Limit" response header`)
	}

	if c.info.Remaining, err = strconv.ParseInt(h.Get("X-Rate-Limit-Remaining"), 10, 64); err != nil {
		return nil, errors.Wrapf(err, `can't parse "X-Rate-Limit-Remaining" response header`)
	}

	reset, err := strconv.ParseInt(h.Get("X-Rate-Reset"), 10, 64)
	if err != nil {
		return nil, errors.Wrapf(err, `can't parse "X-Rate-Reset" response header`)
	}

	c.info.Reset = time.Duration(reset) * time.Second

	if response.StatusCode == http.StatusOK {
		r = &res.Response

		return
	}

	switch response.StatusCode {
	case http.StatusUnauthorized:
		return nil, ErrInvalidAPIKey
	case http.StatusPaymentRequired:
		return nil, ErrSubscriptionIsNotActive
	case http.StatusUnprocessableEntity:
		if res.Error == "Missing 'name' parameter" {
			return nil, ErrMissingName
		}

		return nil, ErrInvalidName
	case http.StatusTooManyRequests:
		if res.Error == "Request limit reached" {
			return nil, ErrRequestLimitReached
		}

		return nil, ErrRequestLimitTooLow
	default:
		return nil, ErrUnknown
	}
}

// Check returns gender info for name.
func (c *Client) Check(name string) (res *Response, err error) {
	u, err := c.url(name)
	if err != nil {
		return nil, errors.Wrap(err, "can't generate API URL")
	}

	if res, err = c.request(u); err != nil {
		return
	}

	return
}