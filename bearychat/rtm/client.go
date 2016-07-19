package rtm

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
)

type client struct {
	api, token string

	c *http.Client
}

func setupClient(api, token string) *client {
	return &client{
		api:   api,
		token: token,
		c:     http.DefaultClient,
	}
}

func (c *client) Get(resource string, out interface{}) (*http.Response, error) {
	return c.Do(resource, "GET", nil, out)
}

func (c *client) Post(resource string, in, out interface{}) (*http.Response, error) {
	return c.Do(resource, "POST", in, out)
}

func (c *client) Do(resource, method string, in, out interface{}) (*http.Response, error) {
	uri, err := url.Parse(fmt.Sprintf("%s/%s", c.api, resource))
	if err != nil {
		return nil, err
	}

	// bind rtm token
	q := uri.Query()
	q.Set("token", c.token)
	uri.RawQuery = q.Encode()

	// build payload
	var buf io.ReadWriter
	if in != nil {
		buf = new(bytes.Buffer)
		err := json.NewEncoder(buf).Encode(in)
		if err != nil {
			return nil, err
		}
	}

	// build request
	req, err := http.NewRequest(method, uri.String(), buf)
	if err != nil {
		return nil, err
	}
	if in != nil {
		req.Header.Set("Content-Type", "application/json")
	}
	req.Header.Set("Accept", "application/json")

	resp, err := c.c.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	resultResponse := new(ResultResponse)
	if err := json.NewDecoder(resp.Body).Decode(resultResponse); err != nil {
		return resp, err
	}

	if resp.StatusCode/100 != 2 || resultResponse.Code != 0 {
		return resp, resultResponse
	}

	if out != nil {
		return resp, json.Unmarshal(resultResponse.Result, out)
	}

	return resp, nil
}

type ResultResponse struct {
	Code        int             `json:"code"`
	Result      json.RawMessage `json:"result"`
	ErrorReason string          `json:"error"`
}

func (r *ResultResponse) Error() string {
	return r.ErrorReason
}
