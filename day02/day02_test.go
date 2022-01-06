package main

import (
	"testing"

	"github.com/matryer/is"
)

func TestParse(t *testing.T) {
	is := is.New(t)
	input := `1-3 a: abcde`
	want := &Password{
		Text:   "abcde",
		Letter: "a",
		Min:    1,
		Max:    3,
	}

	p, err := Parse(input)
	is.NoErr(err)
	is.Equal(p, want)
}

func TestIsValid(t *testing.T) {
	tests := []struct {
		name     string
		password *Password
		want     bool
	}{
		{
			name:     "example 1",
			password: &Password{Text: "abcde", Letter: "a", Min: 1, Max: 3},
			want:     true,
		},
		{
			name:     "example 2",
			password: &Password{Text: "cdefg", Letter: "b", Min: 1, Max: 3},
			want:     false,
		},
		{
			name:     "example 3",
			password: &Password{Text: "ccccccccc", Letter: "c", Min: 2, Max: 9},
			want:     true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			is := is.New(t)
			is.Equal(tt.password.IsValid(), tt.want)
		})
	}
}

func TestIsValidPart2(t *testing.T) {
	tests := []struct {
		name     string
		password *Password
		want     bool
	}{
		{
			name:     "example 1",
			password: &Password{Text: "abcde", Letter: "a", Min: 1, Max: 3},
			want:     true,
		},
		{
			name:     "example 2",
			password: &Password{Text: "cdefg", Letter: "b", Min: 1, Max: 3},
			want:     false,
		},
		{
			name:     "example 3",
			password: &Password{Text: "ccccccccc", Letter: "c", Min: 2, Max: 9},
			want:     false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			is := is.New(t)
			is.Equal(tt.password.IsValidPart2(), tt.want)
		})
	}
}
