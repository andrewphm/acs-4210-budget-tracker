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

func showInputForm(app *tview.Application, state *AppState) {
	form := tview.NewForm().
		AddInputField("Amount", "", 20, nil, func(text string) { state.Amount = text }).
		AddInputField("Description", "", 40, nil, func(text string) { state.Description = text }).
		AddButton("Submit", func() {
			showLoadingScreen(app, state)
		}).
		AddButton("Cancel", func() {
			app.Stop()
			os.Exit(0)
		})
	form.SetBorder(true).SetTitle("Enter Details").SetTitleAlign(tview.AlignLeft)
	app.SetRoot(form, true).SetFocus(form)
}

func getCategoryByTransactionType(transactionType string) []string {
	if transactionType == "Expense" {
			return []string{"Food", "Gifts", "Health/medical", "Home", "Transportation", "Personal", "Pets", "Utilities", "Travel", "Debt", "Other"}
	}
	return []string{"Savings", "Paycheck", "Bonus", "Interest", "Other"}
}

func newCategoryList(app *tview.Application, state *AppState, mainMenu *tview.List) *tview.List {
	list := tview.NewList()
	categories := getCategoryByTransactionType(state.TransactionType)
	for _, category := range categories {
			list.AddItem(category, "", 0, nil)
	}
	list.AddItem("Back", "", 'b', func() {
			app.SetRoot(mainMenu, true)
	})

	list.SetBorder(true).SetTitle("Select category").SetTitleAlign(tview.AlignLeft)
	list.SetSelectedFunc(func(i int, category string, description string, rune rune) {
			state.Category = category
			showInputForm(app, state)
	})
	return list
}



