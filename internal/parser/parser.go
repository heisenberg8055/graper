package http_parser

import (
	"net/http"
	"net/url"

	http_client "github.com/heisenberg8055/graper/internal/api/client"
	http_log "github.com/heisenberg8055/graper/internal/log"
	"golang.org/x/net/html"
)

func ParseHTML(URL string, host string, mp *http_log.Map, parentURL string) {
	ok := mp.Get(URL)
	if ok {
		http_log.LogWarn(http_log.Response{Method: "GET", Url: URL, Status: 0, Message: "Recurred URL"})
		return
	}
	link, err := url.Parse(URL)
	if err != nil {
		http_log.LogErr("URL unable to parse", http_log.Err{Err: err, URL: URL})
		return
	}
	responseChannel := make(chan *http.Response)
	go http_client.RequestURL(URL, responseChannel, mp)
	res := <-responseChannel
	if link.Host == host && len(link.Host) > 0 && res != nil {
		var urls []string = parseBody(res, parentURL)
		for _, url := range urls {
			ParseHTML(url, host, mp, parentURL)
		}
	}
}

func parseBody(body *http.Response, url string) []string {
	var urls []string
	doc, err := html.Parse(body.Body)
	if err != nil {
		http_log.LogErr("Failed to parse data", http_log.Err{Err: err, URL: url})
		return []string{}
	}
	var f func(n *html.Node)
	f = func(n *html.Node) {
		if n.Type == html.ElementNode && n.Data == "a" {
			for _, a := range n.Attr {
				if a.Key == "href" {
					currurl := a.Val
					if len(currurl) > 0 && currurl[0] == '/' {
						urls = append(urls, url+currurl)
					} else {
						urls = append(urls, currurl)
					}
				}
			}
		}
		for c := n.FirstChild; c != nil; c = c.NextSibling {
			f(c)
		}
	}
	f(doc)
	return urls
}
