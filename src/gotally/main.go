package main

import (
	"fmt"
	"gotally/gauth"
	"io/ioutil"
	"log"
	"net/http"

	"golang.org/x/net/context"
	"golang.org/x/oauth2"
	"golang.org/x/oauth2/google"
	"google.golang.org/api/sheets/v4"
)

func getClient(ctx context.Context, config *oauth2.Config) *http.Client {
	// The file token.json stores the user's access and refresh tokens, and is
	// created automatically when the authorization flow completes for the first
	// time.
	tokFile := "token.json"
	tok, err := gauth.TokenFromFile(tokFile)

	if err != nil {
		tok = gauth.GetTokenFromWeb(config)
		gauth.SaveToken(tokFile, tok)
	}

	return config.Client(ctx, tok)
}

func main() {
	ctx := context.Background()
	b, err := ioutil.ReadFile("credentials.json")

	if err != nil {
		log.Fatalf("Unable to read client secret file: %v", err)
	}

	// If modifying these scopes, delete your previously saved token.json.
	config, err := google.ConfigFromJSON(b, "https://www.googleapis.com/auth/spreadsheets")
	if err != nil {
		log.Fatalf("Unable to parse client secret file to config: %v", err)
	}

	client := getClient(ctx, config)
	srv, err := sheets.New(client)

	if err != nil {
		log.Fatalf("Unable to retrieve Sheets client: %v", err)
	}

	s := sheets.SpreadsheetProperties{Title: "My test sheet"}
	rb := &sheets.Spreadsheet{
		Properties: &s,
	}

	resp, err := srv.Spreadsheets.Create(rb).Context(ctx).Do()
	if err != nil {
		log.Fatal(err)
	}

	fmt.Print(resp)
}
