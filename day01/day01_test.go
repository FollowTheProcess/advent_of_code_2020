package main

import (
	"testing"

	"github.com/matryer/is"
)

func TestExamplePart1(t *testing.T) {
	is := is.New(t)
	input := []int{1721, 979, 366, 299, 675, 1456}

	want := 514579

	is.Equal(part1(input), want)
}

func TestExamplePart2(t *testing.T) {
	is := is.New(t)
	input := []int{1721, 979, 366, 299, 675, 1456}

	want := 241861950

	is.Equal(part2(input), want)
}
