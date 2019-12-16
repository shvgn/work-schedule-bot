package spreadsheet

import (
	"fmt"
	"time"
)

// For those who work at night, we count hours before 6AM as previous date
var thresholdHour = 6

// Query presents query for table
type Query struct {
	name      string
	day       string
	sheetName string
}

// NewQuery contructs new query
func NewQuery(name string, t time.Time) *Query {
	tloc := t.Local()

	if tloc.Hour() < thresholdHour {
		tloc = tloc.Add(-24 * time.Hour)
	}

	return &Query{
		name:      name,
		day:       date(tloc),
		sheetName: monthlySheetName(tloc),
	}
}

// DaysRange returns string representation of cells range month days
func (q *Query) DaysRange() string {
	return q.sheetName + "!" + "1:1"
}

// NamesRange returns string representation of cells range of names
func (q *Query) NamesRange() string {
	return q.sheetName + "!" + "A:A"
}

func date(t time.Time) string {
	return fmt.Sprintf("%d", t.Day())
}

func monthlySheetName(t time.Time) string {
	m, y := t.Month(), t.Year()
	return fmt.Sprintf("%s %d", m.String(), y)
}
