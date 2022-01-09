package main

import (
	"testing"

	"github.com/matryer/is"
)

func Test_findSeat(t *testing.T) {
	type args struct {
		code string
	}
	tests := []struct {
		name    string
		args    args
		wantRow int
		wantCol int
		wantErr bool
	}{
		{
			name:    "example 1",
			args:    args{code: "BFFFBBFRRR"},
			wantRow: 70,
			wantCol: 7,
			wantErr: false,
		},
		{
			name:    "example 2",
			args:    args{code: "FFFBBBFRRR"},
			wantRow: 14,
			wantCol: 7,
			wantErr: false,
		},
		{
			name:    "example 3",
			args:    args{code: "BBFFBBFRLL"},
			wantRow: 102,
			wantCol: 4,
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			is := is.New(t)

			gotRow, gotCol, err := findSeat(tt.args.code)

			is.True((err != nil) == tt.wantErr)
			is.Equal(gotRow, tt.wantRow)
			is.Equal(gotCol, tt.wantCol)
		})
	}
}
