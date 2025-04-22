package common

import (
	"bytes"
	"context"
	"crypto/hmac"
	"crypto/sha256"
	"encoding/hex"
	"encoding/json"
	"io"
	"net/http"
	"net/url"
	"strconv"
	"time"
)

// Manages communication with the Duitku API
type ServiceClient struct {
	Cfg *Config
}

// SendAPIRequest sends a request to the Duitku API and returns the response.
//
// It serializes the given request object to JSON and sends it to the API with the
// given method and URL. It then deserializes the response from the server into the
// given response object and returns the response.
//
// If the request or response fails, it will return an error.
func SendAPIRequest(
	ctx context.Context,
	c *ServiceClient,
	req any,
	res any,
	method string,
	url string,
	headerParams map[string]string,
) (*http.Response, error) {
	r, err := c.setRequest(ctx, method, url, req, headerParams)
	if err != nil {
		return nil, err
	}

	httpResp, err := http.DefaultClient.Do(r)
	if err != nil {
		return httpResp, err
	}

	body, err := io.ReadAll(httpResp.Body)
	httpResp.Body.Close()
	httpResp.Body = io.NopCloser(bytes.NewBuffer(body))
	if err != nil {
		return httpResp, err
	}

	if err = json.Unmarshal(body, &res); err != nil {
		return httpResp, err
	}

	return httpResp, nil
}

// setRequest constructs and returns an HTTP request object with the specified
// method, URL, request body, and headers.
//
// It takes the context for the request, the HTTP method, the URL as a string,
// a request body which can be of any type, and a map of header parameters.
// The request body is encoded as JSON if it is not nil.
//
// The function returns a pointer to an http.Request object and an error.
// It returns an error if encoding the request body or creating the request fails.
func (c *ServiceClient) setRequest(
	ctx context.Context,
	method string,
	urlInput string,
	reqBody any,
	headerParams map[string]string,
) (req *http.Request, err error) {
	var body *bytes.Buffer

	if reqBody != nil {
		body = &bytes.Buffer{}
		err = json.NewEncoder(body).Encode(reqBody)
		if err != nil {
			return nil, err
		}
	}

	parsedUrl, err := url.Parse(urlInput)
	if err != nil {
		return nil, err
	}

	if body != nil {
		req, err = http.NewRequestWithContext(ctx, method, parsedUrl.String(), body)
	} else {
		req, err = http.NewRequestWithContext(ctx, method, parsedUrl.String(), nil)
	}
	if err != nil {
		return nil, err
	}

	if len(headerParams) > 0 {
		headers := http.Header{}
		for h, v := range headerParams {
			headers.Set(h, v)
		}
		req.Header = headers
	}

	req.Header.Set("Content-Type", "application/json")

	return req, nil
}

// GetCurrentTimestamp returns the current timestamp in milliseconds
// as a string. It uses the current time in milliseconds since the Unix
// epoch (January 1, 1970 00:00:00 UTC) as the value.
func (c *ServiceClient) GetCurrentTimestamp() string {
	return strconv.FormatInt(time.Now().UnixNano()/1e6, 10)
}

// CreateSignature generates an HMAC SHA-256 signature for the request.
// It combines the merchant code and the provided timestamp, then uses the API key
// as the secret key to create the signature. The result is returned as a
// hexadecimal-encoded string. This signature is used for authenticating requests
// to the Duitku API.
func (c *ServiceClient) CreateSignature(timestamp string) string {
	h := hmac.New(sha256.New, []byte(c.Cfg.APIKey))
	h.Write([]byte(c.Cfg.MerchantCode + timestamp))

	return hex.EncodeToString(h.Sum(nil))
}
