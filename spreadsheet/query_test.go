package spreadsheet

import (
	"reflect"
	"testing"
	"time"
)

func TestNewQuery(t *testing.T) {
	type args struct {
		name string
		t    time.Time
	}

	before6am := time.Date(2019, time.December, 24, 1, 51, 0, 0, time.Local)
	after6am := time.Date(2019, time.December, 24, 11, 51, 0, 0, time.Local)
	before6amDay1 := time.Date(2020, time.January, 1, 1, 51, 0, 0, time.Local)
	after6amDay1 := time.Date(2020, time.January, 1, 11, 51, 0, 0, time.Local)

	tests := []struct {
		name string
		args args
		want *Query
	}{
		{
			"current day before 6AM",
			args{"name", before6am},
			&Query{
				name:      "name",
				day:       "23",
				sheetName: "December 2019",
			},
		},
		{
			"current day after 6AM",
			args{"name", after6am},
			&Query{
				name:      "name",
				day:       "24",
				sheetName: "December 2019",
			},
		},
		{
			"prev month before 6AM on 1st day",
			args{"name", before6amDay1},
			&Query{
				name:      "name",
				day:       "31",
				sheetName: "December 2019",
			},
		},
		{
			"current month after 6AM on 1st day",
			args{"name", after6amDay1},
			&Query{
				name:      "name",
				day:       "1",
				sheetName: "January 2020",
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewQuery(tt.args.name, tt.args.t); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewQuery() = %v, want %v", got, tt.want)
			}
		})
	}
}
