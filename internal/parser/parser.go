package http_parser

import (
	"net/http"
	"net/url"
	"sync"

	http_client "github.com/heisenberg8055/graper/internal/api/client"
	http_log "github.com/heisenberg8055/graper/internal/log"
	"golang.org/x/net/html"
)

func ParseHTML(URL string, host string, mp *http_log.Map, parentURL *url.URL, wg *sync.WaitGroup) {
	defer wg.Done()
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
	responseChannel := make(chan *http.Response, 1)
	go http_client.RequestURL(URL, responseChannel, mp)
	res := <-responseChannel
	if link.Host == host && len(link.Host) > 0 && res != nil {
		var urls []string = parseBody(res, parentURL)
		for _, url := range urls {
			wg.Add(1)
			go ParseHTML(url, host, mp, parentURL, wg)
		}
	}
}

func parseBody(body *http.Response, parent *url.URL) []string {
	var urls []string
	doc, err := html.Parse(body.Body)
	if err != nil {
		http_log.LogErr("Failed to parse data", http_log.Err{
			Err: err,
			URL: parent.String(),
		})
		return []string{}
	}
	var f func(n *html.Node)
	f = func(n *html.Node) {
		if n.Type == html.ElementNode && n.Data == "a" {
			for _, a := range n.Attr {
				if a.Key == "href" {
					currurl := a.Val
					u, err := parent.Parse(currurl)
					if err != nil {
						http_log.LogErr("Bad anchor links", http_log.Err{Err: err, URL: parent.String()})
					} else {
						urls = append(urls, u.String())
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
