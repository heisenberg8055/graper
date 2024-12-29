package tview

import (
	"os/exec"
	"runtime"

	http_log "github.com/heisenberg8055/graper/internal/log"
	"github.com/rivo/tview"
)

func Display(mp *http_log.Map) {
	app := tview.NewApplication()
	list := tview.NewList()
	list.SetTitle("Dead Links")
	for key, val := range mp.Mp {
		if val {
			list.AddItem(key, "", '>', func() {
				var cmd *exec.Cmd

				switch runtime.GOOS {
				case "darwin":
					cmd = exec.Command("open", "-a", "Google Chrome", key)
				case "windows":
					cmd = exec.Command("cmd", "/c", "start", "chrome", key)
				case "linux":
					cmd = exec.Command("xdg-open", key)
				default:

				}
				cmd.Run()
			})
		}
	}
	list.AddItem("Quit", "Press to exit", 'q', func() {
		app.Stop()
	})
	if err := app.SetRoot(list, true).SetFocus(list).Run(); err != nil {
		panic(err)
	}
}
