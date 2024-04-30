package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"

	"golang.org/x/net/context"
	"golang.org/x/oauth2"
	"golang.org/x/oauth2/google"
	"google.golang.org/api/sheets/v4"
	"github.com/rivo/tview"
)

func getClient() *http.Client {
	b, err := ioutil.ReadFile("credentials.json")
	if err != nil {
			log.Fatalf("Unable to read client secret file: %v", err)
	}

	config, err := google.ConfigFromJSON(b, sheets.SpreadsheetsScope)
	if err != nil {
			log.Fatalf("Unable to parse client secret file to config: %v", err)
	}

	tokenFile := "token.json"
	token, err := tokenFromFile(tokenFile)
	if err != nil {
			token = getTokenFromWeb(config)
			saveToken(tokenFile, token)
	}
	return config.Client(context.Background(), token)
}

func getTokenFromWeb(config *oauth2.Config) *oauth2.Token {
	authURL := config.AuthCodeURL("state-token", oauth2.AccessTypeOffline)
	fmt.Printf("Go to the following link in your browser then type the authorization code: \n%v\n", authURL)

	var authCode string
	if _, err := fmt.Scan(&authCode); err != nil {
			log.Fatalf("Unable to read authorization code %v", err)
	}

	token, err := config.Exchange(context.TODO(), authCode)
	if err != nil {
			log.Fatalf("Unable to retrieve token from web %v", err)
	}
	return token
}

func tokenFromFile(file string) (*oauth2.Token, error) {
	f, err := os.Open(file)
	if err != nil {
			return nil, err
	}
	defer f.Close()
	token := &oauth2.Token{}
	err = json.NewDecoder(f).Decode(token)
	return token, err
}

func saveToken(path string, token *oauth2.Token) {
	fmt.Printf("Saving credential file to: %s\n", path)
	f, err := os.OpenFile(path, os.O_RDWR|os.O_CREATE|os.O_TRUNC, 0600)
	if err != nil {
			log.Fatalf("Unable to cache oauth token: %v", err)
	}
	defer f.Close()
	json.NewEncoder(f).Encode(token)
}

func writeToSheet(app *tview.Application, state *AppState) {
	client := getClient()

	srv, err := sheets.New(client)
	if err != nil {
			log.Printf("Unable to retrieve Sheets client: %v", err)
			return
	}

	var writeRange string
	if state.TransactionType == "Expense" {
			writeRange = "Transactions!B6:E6"
	} else {
			writeRange = "Transactions!G6:J6"
	}

	spreadsheetId := "1xmbTeGZn6zJcFFjcq2Rm63Bh7IQvn8R7YeX6RjJuMiY"
	vr := &sheets.ValueRange{
			Values: [][]interface{}{{state.Date, state.Amount, state.Description, state.Category}},
	}

	_, err = srv.Spreadsheets.Values.Append(spreadsheetId, writeRange, vr).ValueInputOption("USER_ENTERED").InsertDataOption("INSERT_ROWS").Do()
	if err != nil {
			state.StatusMessage = fmt.Sprintf("Failed to write data: %v", err)
	} else {
			state.StatusMessage = "Data written successfully!"
			fmt.Println("Data written successfully!")
	}
}
