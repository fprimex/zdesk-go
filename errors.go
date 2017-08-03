package zdesk

import (
	"bytes"
	"fmt"
	"net/http"
)

// Ensure HTTPError is, in fact, an error.
var _ error = (*HTTPError)(nil)

// HTTPError is a custom error type that wraps an HTTP status code with some
// helper functions.
type HTTPError struct {
	// StatusCode is the HTTP status code (2xx-5xx).
	StatusCode int

	// Message and Detail are information returned by the API.
	Message string
	Detail  string
}

// NewHTTPError creates a new HTTP error from the given code.
func NewHTTPError(resp *http.Response) *HTTPError {
	var e HTTPError
	//if resp.Body != nil {
	//	e.Message = resp.Body
	//}
	e.StatusCode = resp.StatusCode
	return &e
}

// Error implements the error interface and returns the string representing the
// error text that includes the status code and the corresponding status text.
func (e *HTTPError) Error() string {
	var r bytes.Buffer
	fmt.Fprintf(&r, "%d - %s", e.StatusCode, http.StatusText(e.StatusCode))

	if e.Message != "" {
		fmt.Fprintf(&r, "\nMessage: %s", e.Message)
	}

	if e.Detail != "" {
		fmt.Fprintf(&r, "\nDetail: %s", e.Detail)
	}

	return r.String()
}

// String implements the stringer interface and returns the string representing
// the string text that includes the status code and corresponding status text.
func (e *HTTPError) String() string {
	return e.Error()
}

// IsNotFound returns true if the HTTP error code is a 404, false otherwise.
func (e *HTTPError) IsNotFound() bool {
	return e.StatusCode == 404
}
