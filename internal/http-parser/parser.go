package http_parser

import (
	"net/http"
	"net/url"
	"sync"
	"time"

	http_client "github.com/heisenberg8055/graper/internal/api/client"
	http_log "github.com/heisenberg8055/graper/internal/log"
)

type Map struct {
	mp map[string]*http.Response
	mu *sync.Mutex
}

func newMap() *Map {
	return &Map{mp: make(map[string]*http.Response)}
}

func (m *Map) Set(key string, value *http.Response) {
	m.mu.Lock()
	defer m.mu.Unlock()
	m.mp[key] = value
}

func (m *Map) Get(key string) bool {
	m.mu.Lock()
	defer m.mu.Unlock()
	_, ok := m.mp[key]
	return ok
}

func Crawler(isRecur bool, arg string) {
	link, err := url.Parse(arg)
	if err != nil {
		http_log.LogErr(http_log.Err{Err: err, URL: arg})
		return
	}
	mp := newMap()
	if !isRecur {
		parseSingleHTML(arg)
	} else {
		parseHTML(arg, mp, link.Host)
	}
}

func parseSingleHTML(url string) {
	t := time.Now()
	body, err := http_client.GetBody(url)
	elapsed := time.Since(t).String()
	if err != nil {
		http_log.LogError(http_log.Response{Method: "GET", Time_taken: elapsed, Url: url, Status: 0, Message: err.Error()})
		return
	}
	if body.StatusCode >= 400 {
		http_log.LogError(http_log.Response{Method: "GET", Time_taken: elapsed, Url: url, Status: body.StatusCode, Message: "Dead Link"})
		return
	}
	if body.StatusCode >= 299 && body.StatusCode <= 399 {
		http_log.LogWarn(http_log.Response{Method: "GET", Time_taken: elapsed, Url: url, Status: body.StatusCode, Message: "Redirect"})
		return
	}
	http_log.LogInfo(http_log.Response{Method: "GET", Time_taken: elapsed, Url: url, Status: body.StatusCode, Message: "Success"})
}

func parseHTML(url string, mp *Map, host string) {

}
