package spreadsheet

import (
	"fmt"
	"strings"
)

// RecordOpts to have named opts to create Record
type RecordOpts struct {
	sheetName string
	row       string
	startCol  string
	endCol    string
}

// Record of time in spreadsheet
type Record struct {
	sheetName string
	row       string
	startCol  string
	endCol    string
	Starts    []string
	Ends      []string
}

// NewRecord creates range struct
func NewRecord(ro *RecordOpts) *Record {
	return &Record{
		ro.sheetName,
		ro.row,
		ro.startCol,
		ro.endCol,
		[]string{},
		[]string{},
	}
}

// Range returns full cells range in spreadsheet
func (r *Record) Range() string {
	cellRange := fmt.Sprintf("%s%s:%s%s", r.startCol, r.row, r.endCol, r.row)
	return r.sheetName + "!" + cellRange
}

// Payload for spreadsheet API
func (r *Record) Payload() [][]interface{} {
	return [][]interface{}{{
		strings.Join(r.Starts, "\n"),
		strings.Join(r.Ends, "\n"),
	}}
}

// Push adds value to starts or ends depending on the content
func (r *Record) Push(s string) {
	if len(r.Starts) > len(r.Ends) {
		r.Ends = append(r.Ends, s)
		return
	}
	r.Starts = append(r.Starts, s)
}
