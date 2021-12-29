package gohttp

import (
	"errors"
	"net/http"
)

func (c *httpClient) do(method, url string, headers http.Header, body interface{}) (*http.Response, error) {
	client := http.Client{}

	request, err := http.NewRequest(method, url, nil)
	if err != nil {
		return nil, errors.New("Unable to create a new request")
	}

	return client.Do(request)
}
