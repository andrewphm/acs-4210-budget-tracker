package main

import (
	"os"
	"github.com/rivo/tview"
)

func setupAppUI(app *tview.Application, state *AppState) {
	var mainMenu *tview.List
	mainMenu = tview.NewList().
		AddItem("Expense", "", 'e', func() {
			state.TransactionType = "Expense"
			app.SetRoot(newCategoryList(app, state, mainMenu), true)
		}).
		AddItem("Income", "", 'i', func() {
			state.TransactionType = "Income"
			app.SetRoot(newCategoryList(app, state, mainMenu), true)
		}).
		AddItem("Quit", "Press to exit", 'q', func() {
			app.Stop()
			os.Exit(0)
		})

	mainMenu.SetBorder(true).SetTitle("Select Transaction Type").SetTitleAlign(tview.AlignLeft)

	app.SetRoot(mainMenu, true).SetFocus(mainMenu)
}

func showLoadingScreen(app *tview.Application, state *AppState) {
	loadingScreen := tview.NewModal().SetText("Finished!\n" + state.TransactionType + "-" + state.Category + "\n" + state.Description + "\n" + "$" + state.Amount + "\n" + state.Date)
	app.SetRoot(loadingScreen, true)
	writeToSheet(app, state)
}

func showFinishedScreen(app *tview.Application, message string) {
	finishedScreen := tview.NewModal().SetText("DOne!")
	app.SetRoot(finishedScreen, true)
}


