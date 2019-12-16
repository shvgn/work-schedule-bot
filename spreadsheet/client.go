package spreadsheet

import (
	"errors"
	"fmt"
	"regexp"
	"strings"

	"google.golang.org/api/sheets/v4"
)

// Client provides CRUD for spreadsheet
type Client struct {
	service       *sheets.Service
	spreadsheetID string
}

// NewClient inits new spreadsheet client
func NewClient(credsPath string, spreadsheetID string) (*Client, error) {
	service, err := getSheetService(credsPath)
	if err != nil {
		return nil, err
	}
	return &Client{service, spreadsheetID}, nil
}

// Set record value to spreadsheet
func (c *Client) Set(r *Record) error {
	_, err := c.service.Spreadsheets.Values.Update(
		c.spreadsheetID,
		r.Range(),
		&sheets.ValueRange{
			Values:         r.Payload(),
			MajorDimension: "ROWS",
		},
	).ValueInputOption("USER_ENTERED").Do()

	if err != nil {
		return fmt.Errorf("cannot set value: %v", err)
	}
	return nil
}

// Clear the record
func (c *Client) Clear(q *Query) error {
	rec, err := c.findRecord(q)
	if err != nil {
		return err
	}
	_, err = c.service.Spreadsheets.Values.Clear(
		c.spreadsheetID,
		rec.Range(),
		&sheets.ClearValuesRequest{}).Do()

	if err != nil {
		return fmt.Errorf("cannot clear cells: %v", err)
	}
	return nil
}

// Get retrieves record from spreadsheet
func (c *Client) Get(q *Query) (*Record, error) {
	rec, err := c.findRecord(q)
	if err != nil {
		return nil, fmt.Errorf("cannot find table range: %v", err)
	}

	resp, err := c.service.Spreadsheets.Values.Get(c.spreadsheetID, rec.Range()).Do()

	if err != nil {
		return nil, fmt.Errorf("cannot read time: %v", err)
	}

	if len(resp.Values) == 0 {
		// No data found
		return rec, nil
	}

	row := resp.Values[0]
	if len(row) > 2 {
		return nil, fmt.Errorf("unaccepted response data: %v", row)
	}

	if len(row) > 0 {
		cellText := row[0].(string)
		rec.Starts = strings.Split(cellText, "\n")
	}
	if len(row) > 1 {
		cellText := row[1].(string)
		rec.Ends = strings.Split(cellText, "\n")
	}

	return rec, nil
}

func (c *Client) findRecord(q *Query) (*Record, error) {
	// TODO perf: find row and colums with two goroutines or make single request to the
	// table for the whole value body
	row, err := c.findUserRow(q)
	if err != nil {
		return nil, fmt.Errorf("cannot find record row: %v", err)
	}

	startCol, endCol, err := c.findDayColumns(q)
	if err != nil {
		return nil, fmt.Errorf("cannot find record columns: %v", err)
	}

	rec := NewRecord(&RecordOpts{
		sheetName: q.sheetName,
		row:       row,
		startCol:  startCol,
		endCol:    endCol,
	})
	return rec, nil
}

func (c *Client) findUserRow(q *Query) (string, error) {
	resp, err := c.service.Spreadsheets.Values.Get(c.spreadsheetID, q.NamesRange()).Do()
	if err != nil {
		return "", err
	}

	re := regexp.MustCompile(q.name)
	for i, row := range resp.Values {
		if len(row) < 1 {
			continue
		}
		value := row[0].(string)
		if !re.MatchString(value) {
			continue
		}
		return fmt.Sprintf("%d", i+1), nil
	}
	return "", fmt.Errorf("did not find \"%s\" among names", q.name)
}

func (c *Client) findDayColumns(q *Query) (string, string, error) {
	resp, err := c.service.Spreadsheets.Values.Get(c.spreadsheetID, q.DaysRange()).Do()
	if err != nil {
		return "", "", err
	}

	if len(resp.Values) == 0 {
		return "", "", errors.New("could not fetch days row")
	}

	// First column is for names, then come three columns per day (start, end,
	// and extra), thus at max we have 1+3×31 = 94. 'DA' is 26×4 = 104th. The
	// last I saw in december was CP, so we assume extra columns here.
	minimalMonthCells := 28*3 + 1 // This is a naive asertion, we take February as the shortest to campare with
	row := resp.Values[0]

	if len(row) < minimalMonthCells {
		err := fmt.Errorf(
			"cannot believe we don't have enough cols: %d, expected %d at minimum",
			len(resp.Values), minimalMonthCells)
		return "", "", err
	}

	// Find columns of the start and the end of the day
	for i, v := range row {
		if q.day != v.(string) {
			continue
		}

		startCol, err := colAddrByIndex(i)
		if err != nil {
			return "", "", err
		}

		endCol, err := colAddrByIndex(i + 1)
		if err != nil {
			return "", "", err
		}

		return startCol, endCol, nil
	}

	return "", "", fmt.Errorf("could not find column for day %s", q.day)
}
