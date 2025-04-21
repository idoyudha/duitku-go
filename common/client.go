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

func SendAPIRequest(
	ctx context.Context,
	c *ServiceClient,
	req any,
	res any,
	method string,
	path string,
	headerParams map[string]string,
) (*http.Response, error) {
	r, err := c.setRequest(ctx, method, path, req, headerParams)
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

func (c *ServiceClient) setRequest(
	ctx context.Context,
	method string,
	path string,
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

	baseUrl := SandboxBaseURL
	if c.Cfg.Environment == ProductionEnv {
		baseUrl = ProductionBaseURL
	}

	url, err := url.Parse(baseUrl + path)
	if err != nil {
		return nil, err
	}

	if body != nil {
		req, err = http.NewRequestWithContext(ctx, method, url.String(), body)
	} else {
		req, err = http.NewRequestWithContext(ctx, method, url.String(), nil)
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

func (c *ServiceClient) GetCurrentTimestamp() string {
	return strconv.FormatInt(time.Now().UnixNano()/1e6, 10)
}

func (c *ServiceClient) CreateSignature(timestamp string) string {
	h := hmac.New(sha256.New, []byte(c.Cfg.APIKey))
	h.Write([]byte(c.Cfg.MerchantCode + timestamp))

	return hex.EncodeToString(h.Sum(nil))
}
