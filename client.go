package genderize

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
	"time"
)

// Info API rate limits info.
type Info struct {
	Limit     int64
	Remaining int64
	Reset     time.Duration
}

// Error response error.
type Error struct {
	Error string `json:"error,omitempty"`
}

// Client genderize API client.
type Client struct {
	options *Options
}

// Execute executes API request and returns result.
func (c *Client) Execute(request *Request) (collection *Collection, err error) {
	u := request.Encode(c.options.APIKey)

	req, err := http.NewRequestWithContext(request.ctx, http.MethodGet, u, nil)
	if err != nil {
		return
	}

	res, err := c.options.HTTPClient.Do(req)
	if err != nil {
		return
	}

	defer func() {
		_ = res.Body.Close()
	}()

	collection, err = c.processAPIResponse(res)

	return
}

// ExecuteX like Execute, but panics when error.
func (c *Client) ExecuteX(request *Request) *Collection {
	collection, err := c.Execute(request)
	if err != nil {
		panic(err)
	}

	return collection
}

func (c *Client) processAPIResponse(res *http.Response) (collection *Collection, err error) {
	collection = &Collection{}

	collection.info, err = c.processInfo(res)
	if err != nil {
		return
	}

	switch res.StatusCode {
	case http.StatusOK:
		collection.genders, err = c.processResponse(res)
	case http.StatusUnauthorized:
		err = ErrInvalidAPIKey
	case http.StatusPaymentRequired:
		err = ErrSubscriptionIsNotActive
	case http.StatusUnprocessableEntity:
		err = fmt.Errorf(`%w cause %s`, ErrValidation, c.processError(res))
	case http.StatusTooManyRequests:
		err = fmt.Errorf(`%w cause %s`, ErrTooManyRequests, c.processError(res))
	default:
		err = ErrInternal
	}

	return
}

func (c *Client) processError(res *http.Response) (s string) {
	var e Error
	_ = json.NewDecoder(res.Body).Decode(&e)

	s = e.Error

	return
}

func (c *Client) processInfo(res *http.Response) (info *Info, err error) {
	i := &Info{}

	if i.Limit, err = c.processHeader(res, HdrXRateLimitLimit); err != nil {
		err = fmt.Errorf(`%w %s`, ErrResponseHeader, err)

		return
	}

	if i.Remaining, err = c.processHeader(res, HdrXRateLimitRemaining); err != nil {
		return
	}

	reset, err := c.processHeader(res, HdrXRateReset)
	if err == nil {
		i.Reset = time.Duration(reset) * time.Second
	}

	info = i

	return
}

func (c *Client) processHeader(res *http.Response, header string) (value int64, err error) {
	v := res.Header.Get(header)

	value, err = strconv.ParseInt(v, 10, 64)
	if err != nil {
		err = fmt.Errorf(`(%s: %s): %w`, header, v, err)
	}

	return
}

func (c *Client) processResponse(res *http.Response) (genders map[string]*Gender, err error) {
	var apiRes []*Gender
	if err = json.NewDecoder(res.Body).Decode(&apiRes); err != nil {
		err = fmt.Errorf("%w: %s", ErrResponseBody, err)

		return
	}

	genders = map[string]*Gender{}

	for _, g := range apiRes {
		genders[g.Name] = g
	}

	return
}

// NewClient returns new API client instance.
func NewClient(options ...Option) *Client {
	client := &Client{
		options: &Options{
			HTTPClient: http.DefaultClient,
		},
	}

	for _, opt := range options {
		opt(client.options)
	}

	return client
}
