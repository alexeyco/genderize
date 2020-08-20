package genderize

import (
	"context"
	"encoding/json"
	"net/http"
	"strconv"
	"time"
)

// Client genderize API client.
type Client struct {
	options *Options
}

// Execute executes API request and returns result.
func (c *Client) Execute(ctx context.Context, request *Request) (response *Response, info *Info, err error) {
	req, err := http.NewRequestWithContext(ctx, http.MethodGet, request.Encode(c.options.APIKey), nil)
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

	response, info, err = c.processAPIResponse(res)

	return
}

func (c *Client) processAPIResponse(res *http.Response) (response *Response, info *Info, err error) {
	var apiRes apiResponse
	if err = json.NewDecoder(res.Body).Decode(&apiRes); err != nil {
		return
	}

	if res.StatusCode == http.StatusOK {
		if info, err = c.processInfo(res); err != nil {
			return
		}

		response, err = c.processResponse(res)

		return
	}

	return
}

func (c *Client) processInfo(res *http.Response) (info *Info, err error) {
	i := &Info{}

	if i.Limit, err = c.processHeader(res, "X-Rate-Limit-Limit"); err != nil {
		return
	}

	if i.Remaining, err = c.processHeader(res, "X-Rate-Limit-Remaining"); err != nil {
		return
	}

	reset, err := c.processHeader(res, "X-Rate-Reset")
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
		err = &ErrResponseHeader{
			header: header,
			value:  v,
			err:    err,
		}
	}

	return
}

func (c *Client) processResponse(res *http.Response) (response *Response, err error) {
	var apiRes apiResponse
	if err = json.NewDecoder(res.Body).Decode(&apiRes); err != nil {
		err = &ErrResponse{
			err: err,
		}

		return
	}

	response = &apiRes.Response

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
