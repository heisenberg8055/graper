package http_client

import (
	"net/http"
	"time"

	http_log "github.com/heisenberg8055/graper/internal/log"
)

func newClient() *http.Client {
	c := http.Client{
		CheckRedirect: func(req *http.Request, via []*http.Request) error {
			// Prevent following the redirect
			return http.ErrUseLastResponse
		},
		Timeout: time.Second * 5,
	}
	return &c
}

func RequestURL(url string, response chan<- *http.Response, mp *http_log.Map) {
	c := newClient()
	req, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		http_log.LogError(http_log.Response{Method: "GET", Url: url, Status: 0, Message: err.Error()})
		close(response)
		return
	}
	mp.Set(url, false)
	res, err := c.Do(req)
	if err != nil {
		http_log.LogError(http_log.Response{Method: "GET", Url: url, Status: 0, Message: err.Error()})
		mp.Set(url, true)
		response <- res
		return
	}
	if res.StatusCode >= 400 {
		http_log.LogError(http_log.Response{Method: "GET", Url: url, Status: res.StatusCode, Message: "Dead Link"})
		mp.Set(url, true)
		response <- res
		return
	}
	if res.StatusCode >= 300 && res.StatusCode <= 399 {
		http_log.LogWarn(http_log.Response{Method: "GET", Url: url, Status: res.StatusCode, Message: "Redirect"})
		response <- res
		return
	}
	if res.StatusCode >= 200 && res.StatusCode <= 299 {
		http_log.LogInfo(http_log.Response{Method: "GET", Url: url, Status: res.StatusCode, Message: "Success"})
		response <- res
		return
	} else {
		http_log.LogError(http_log.Response{Method: "GET", Url: url, Status: http.StatusInternalServerError, Message: "Server Error"})
		mp.Set(url, true)
		close(response)
		return
	}
}
