package testhelpers

import (
	"net/http"
)

type HttpClientMock struct {
	Transport http.RoundTripper
}

// Implement the Do method of the mock HTTP client
func (c *HttpClientMock) Do(req *http.Request) (*http.Response, error) {
	return c.Transport.RoundTrip(req)
}

