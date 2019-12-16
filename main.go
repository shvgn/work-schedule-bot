package main

import (
	"io/ioutil"
	"log"
	"os"

	"github.com/shvgn/work-schedule-bot/bot"
	"github.com/shvgn/work-schedule-bot/spreadsheet"
	"github.com/shvgn/work-schedule-bot/table"
)

func main() {
	spreadsheetID, ok := os.LookupEnv("SPREADSHEET_ID")
	if !ok {
		log.Fatal("SPREADSHEET_ID is not provided in env")
	}

	client, err := spreadsheet.NewClient("secrets/credentials.json", spreadsheetID)
	if err != nil {
		log.Fatalf("cannot access spreadsheet: %v", err)
	}

	tbl := table.NewTable(client)

	token, err := ioutil.ReadFile("secrets/telegram_token")
	if err != nil {
		log.Fatalf("unable to read telegram token file: %v", err)
	}

	b, err := bot.InitBot(string(token), tbl)
	if err != nil {
		log.Fatalf("unable to init telegram bot: %v", err)
	}

	b.Start()
}
