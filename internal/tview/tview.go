package tview

import (
	"os/exec"
	"runtime"

	http_log "github.com/heisenberg8055/graper/internal/log"
	"github.com/rivo/tview"
)

func Display(mp *http_log.Map, url string) {
	app := tview.NewApplication()
	list := tview.NewList()
	list.SetTitle("Dead Links")
	var brokenLinks []string
	for key, val := range mp.Mp {
		if val {
			brokenLinks = append(brokenLinks, key)
		}
	}
	for _, val := range brokenLinks {
		list.AddItem(val, "", '>', func() {
			var cmd *exec.Cmd

			switch runtime.GOOS {
			case "darwin":
				cmd = exec.Command("open", "-a", "Google Chrome", val)
			case "windows":
				cmd = exec.Command("cmd", "/c", "start", "chrome", val)
			case "linux":
				cmd = exec.Command("xdg-open", val)
			default:

			}
			cmd.Run()
		})
	}
	if len(brokenLinks) == 0 {
		list.AddItem(url, "No Broken Links", '>', func() {
			var cmd *exec.Cmd

			switch runtime.GOOS {
			case "darwin":
				cmd = exec.Command("open", "-a", "Google Chrome", url)
			case "windows":
				cmd = exec.Command("cmd", "/c", "start", "chrome", url)
			case "linux":
				cmd = exec.Command("xdg-open", url)
			default:

			}
			cmd.Run()
		})
	}
	list.AddItem("Quit", "Press to exit", 'q', func() {
		app.Stop()
	})
	if err := app.SetRoot(list, true).SetFocus(list).Run(); err != nil {
		panic(err)
	}
}
