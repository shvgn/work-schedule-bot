package table

import (
	"fmt"
	"time"

	"github.com/shvgn/work-schedule-bot/spreadsheet"
)

// Table represents the time table for editing
type Table struct {
	client *spreadsheet.Client
}

// NewTable creates new table for handling spreadsheet
func NewTable(c *spreadsheet.Client) *Table {
	return &Table{c}
}

// Append time record for user
func (tbl *Table) Append(name string, t time.Time) (*spreadsheet.Record, error) {
	tround := t.Round(5 * time.Minute)

	q := spreadsheet.NewQuery(name, tround)
	record, err := tbl.client.Get(q)

	if err != nil {
		return nil, err
	}

	value := tround.Local().Format("15:04")
	record.Push(value)
	err = tbl.client.Set(record)

	if err != nil {
		return nil, err
	}
	return record, nil
}

// Clear deletes data for today
func (tbl *Table) Clear(name string, t time.Time) error {
	tround := t.Round(5 * time.Minute)
	q := spreadsheet.NewQuery(name, tround)
	err := tbl.client.Clear(q)
	if err != nil {
		return fmt.Errorf("cannot clear: %v", err)
	}
	return nil
}
