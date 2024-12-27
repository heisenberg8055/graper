package http_client

import (
	"net/http"
	"time"
)

func newClient() *http.Client {
	c := http.Client{CheckRedirect: func(req *http.Request, via []*http.Request) error {
		return http.ErrUseLastResponse
	},
		Timeout: time.Second * 5}
	return &c
}

func GetBody(URL string) (*http.Response, error) {
	c := newClient()
	req, err := http.NewRequest(http.MethodGet, URL, nil)
	if err != nil {
		return nil, err
	}
	return c.Do(req)
}
