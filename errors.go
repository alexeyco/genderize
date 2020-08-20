package genderize

import "fmt"

// ErrResponseHeader API response header error.
type ErrResponseHeader struct {
	header string
	value  string
	err    error
}

// Header returns problem header.
func (e *ErrResponseHeader) Header() string {
	return e.header
}

// Value returns problem header value.
func (e *ErrResponseHeader) Value() string {
	return e.value
}

// Error returns error as a string.
func (e *ErrResponseHeader) Error() string {
	return fmt.Sprintf(`response header "%s" with value "%s" is wrong cause %s`, e.header, e.value, e.err)
}

// ErrResponse response error.
type ErrResponse struct {
	body []byte
	err  error
}

// Body returns response body as a string.
func (e *ErrResponse) Body() string {
	return string(e.body)
}

// Error returns error as a string.
func (e *ErrResponse) Error() string {
	return fmt.Sprintf(`response body "%s" is wronf cause %s`, string(e.body), e.err)
}
