package main

import (
	"github.com/rivo/tview"
	"time"
)

func getCurrentFormattedDate() string {
	return time.Now().Format("01/02/2006")
}

func main() {
    client := getClient()

    if client != nil {
        app := tview.NewApplication()
        state := &AppState{
                Date: getCurrentFormattedDate(),
            }
        setupAppUI(app, state)
        if err := app.Run(); err != nil {
            panic(err)
        }
    }
}
