package crawler

import (
	"net/url"

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
	http_parser.ParseHTML(url, "", mp, url)
	tview.Display(mp)
}

func parseRecurRequests(url, host string) {
	mp := http_log.NewMap()
	http_parser.ParseHTML(url, host, mp, url)
	tview.Display(mp)
}
