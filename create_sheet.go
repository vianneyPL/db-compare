package main

// BEFORE RUNNING:
// ---------------
// 1. If not already done, enable the Google Sheets API
//    and check the quota for your project at
//    https://console.developers.google.com/apis/api/sheets
// 2. Install and update the Go dependencies by running `go get -u` in the
//    project directory.

import (
	"log"
	"net/http"

	"golang.org/x/net/context"
	sheets "google.golang.org/api/sheets/v4"
)

// CreateSheet : creates a google spreadsheet and returns its id and url
func CreateSheet(client *http.Client) (string, string) {
	ctx := context.Background()
	sheetsService, err := sheets.New(client)
	if err != nil {
		log.Fatal(err)
	}

	rb := &sheets.Spreadsheet{
	// TODO: Add desired fields of the request body.
	}

	resp, err := sheetsService.Spreadsheets.Create(rb).Context(ctx).Do()
	if err != nil {
		log.Fatal(err)
	}
	return resp.SpreadsheetId, resp.SpreadsheetUrl
}
