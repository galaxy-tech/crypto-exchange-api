package http

import (
	"io/ioutil"
	"net/http"
)

type HTTPClient struct {
	client *http.Client
}

func NewHTTPClient() *HTTPClient {
	return &HTTPClient{client: http.DefaultClient}
}

func (httpClient *HTTPClient) GetHTTP(url string) ([]byte, error) {
	method := "GET"
	req, err := http.NewRequest(method, url, nil)

	if err != nil {
		return nil, err
	}
	res, err := httpClient.client.Do(req)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()

	return ioutil.ReadAll(res.Body)
}