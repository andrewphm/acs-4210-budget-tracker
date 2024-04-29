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
