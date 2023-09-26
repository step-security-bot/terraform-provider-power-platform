package clients

import (
	"net/http"
)

type BapiRoundTripper struct {
	Transport http.RoundTripper
}

func (b *BapiRoundTripper) RoundTrip(req *http.Request) (*http.Response, error) {

	if req.Header.Get("Content-Type") == "" {
		req.Header.Set("Content-Type", "application/json")
	}

	req.Header.Set("User-Agent", "terraform-provider-power-platform")

	r, err := b.Transport.RoundTrip(req)
	if err != nil {
		return nil, err
	}

	if r.StatusCode != http.StatusOK {
		return r, err
	}

	return r, nil
}
