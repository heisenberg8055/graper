package crawler

import (
	"net/url"
	"sync"

	http_log "github.com/heisenberg8055/graper/internal/log"
	http_parser "github.com/heisenberg8055/graper/internal/parser"
	"github.com/heisenberg8055/graper/internal/tview"
)

func Crawler(isRecur bool, arg string) {
	if !isRecur {
		parseSingleRequest(arg)
	} else {
		link, err := url.Parse(arg)
		if err != nil {
			panic(err)
		}
		parseRecurRequests(arg, link.Host)
	}
}

func parseSingleRequest(url string) {
	mp := http_log.NewMap()
	var wg sync.WaitGroup
	wg.Add(1)
	go http_parser.ParseHTML(url, "", mp, nil, &wg)
	wg.Wait()
	tview.Display(mp, url)
}

func parseRecurRequests(arg, host string) {
	mp := http_log.NewMap()
	parent, err := url.Parse(arg)
	if err != nil {
		http_log.LogErr("Unable to parse input URL", http_log.Err{Err: err, URL: arg})
	}
	var wg sync.WaitGroup
	wg.Add(1)
	go http_parser.ParseHTML(arg, host, mp, parent, &wg)
	wg.Wait()
	tview.Display(mp, arg)
}
