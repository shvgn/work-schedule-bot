package spreadsheet

import (
	"testing"
)

func Test_letterByIndex(t *testing.T) {
	type args struct {
		i int
	}
	tests := []struct {
		name    string
		args    args
		want    string
		wantErr bool
	}{
		{"0 is A", args{i: 0}, "A", false},
		{"25 is Z", args{i: 25}, "Z", false},
		{"-1 is error", args{i: -1}, "", true},
		{"-2 is error", args{i: -2}, "", true},
		{"26 is error", args{i: 26}, "", true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := letterByIndex(tt.args.i)
			if (err != nil) != tt.wantErr {
				t.Errorf("letterByIndex() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("letterByIndex() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_colAddrByIndex(t *testing.T) {
	type args struct {
		i int
	}
	tests := []struct {
		name    string
		args    args
		want    string
		wantErr bool
	}{
		{"-1 is error", args{i: -1}, "", true},
		{"-2 is error", args{i: -2}, "", true},
		{"0 is A", args{i: 0}, "A", false},
		{"1 is B", args{i: 1}, "B", false},
		{"2 is C", args{i: 2}, "C", false},
		{"24 is Y", args{i: 24}, "Y", false},
		{"25 is Z", args{i: 25}, "Z", false},
		{"26 is AA", args{i: 26}, "AA", false},
		{"27 is AB", args{i: 27}, "AB", false},
		{"28 is AC", args{i: 28}, "AC", false},
		{"50 is AY", args{i: 50}, "AY", false},
		{"51 is AZ", args{i: 51}, "AZ", false},
		{"52 is BA", args{i: 52}, "BA", false},
		{"53 is BB", args{i: 53}, "BB", false},
		{"76 is BY", args{i: 76}, "BY", false},
		{"77 is BZ", args{i: 77}, "BZ", false},
		{"93 is CP", args{i: 93}, "CP", false},
		{"26×26+25 is ZZ", args{i: 26*26 + 25}, "ZZ", false},
		{"26×27 is AAA", args{i: 26 * 27}, "AAA", false},
		{"26×27+1 is AAB", args{i: 26*27 + 1}, "AAB", false},
		{"26×27+25 is AAZ", args{i: 26*27 + 25}, "AAZ", false},
		{"26×28 is ABA", args{i: 26 * 28}, "ABA", false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := colAddrByIndex(tt.args.i)
			if (err != nil) != tt.wantErr {
				t.Errorf("colAddrByIndex() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("colAddrByIndex() = %v, want %v", got, tt.want)
			}
		})
	}
}
