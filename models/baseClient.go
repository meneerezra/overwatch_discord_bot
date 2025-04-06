package main

import (
	"encoding/json"
	"errors"
	"io"
	"log"
	"net/http"
	"net/url"
	"strings"
)

var (
	ErrorInvalidRequest  = errors.New("invalid request")
	ErrorInvalidURL      = errors.New("invalid URL")
	ErrorCannotParsePath = errors.New("cannot parse path")
	ErrorCannotReadBody  = errors.New("Cannot read body")
	Error401             = errors.New("Unauthorized")
	Error403             = errors.New("Forbidden")
	Error404             = errors.New("Not Found")
	Error500             = errors.New("Internal Server Error")
)

type Client struct {
	client  *http.Client
	BaseURL string
}

func (client *Client) newRequest(method, path string, headers http.Header, params url.Values) (*http.Request, error) {
	if params == nil {
		params = url.Values{}
	}
	queryString := params.Encode()
	var body io.Reader
	fullURL := client.BaseURL + path

	if queryString != "" && method == http.MethodGet {
		fullURL += "?" + queryString
	}

	if method != http.MethodGet {
		body = strings.NewReader(queryString)
	}

	req, err := http.NewRequest(method, fullURL, body)
	if err != nil {
		return nil, err
	}

	if method == http.MethodPost || method == http.MethodPut {
		req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	}
	for k, v := range headers {
		req.Header[k] = v
	}
	return req, nil
}

func (client *Client) do(req *http.Request, v interface{}) error {
	var (
		err     error
		resp    *http.Response
		retries = 3
	)

	for retries > 0 {
		resp, err = client.client.Do(req)
		if shouldRetry(err, resp) {
			log.Println(err)
			retries -= 1
			continue
		}

		if err != nil {
			return err
		}

		defer resp.Body.Close()

		break
	}

	if resp != nil {
		switch resp.StatusCode {
		case http.StatusNotFound:
			return Error404
		case http.StatusUnauthorized:
			return Error401
		case http.StatusForbidden:
			return Error403
		case http.StatusInternalServerError:
			return Error500
		}

		defer resp.Body.Close()

		if v != nil {
			err = json.NewDecoder(resp.Body).Decode(v)
			if err != nil {
				return err
			}
		}
	}
	return err
}

func shouldRetry(err error, resp *http.Response) bool {
	if err != nil {
		return true
	}
	if resp.StatusCode == http.StatusBadGateway ||
		resp.StatusCode == http.StatusServiceUnavailable ||
		resp.StatusCode == http.StatusGatewayTimeout {
		return true
	}
	return false
}
