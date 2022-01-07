package main

import "testing"

func TestExamplePart1(t *testing.T) {
	input := []byte(`..##.......
#...#...#..
.#....#..#.
..#.#...#.#
.#...##..#.
..#.##.....
.#.#.#....#
.#........#
#.##...#...
#...##....#
.#..#...#.#`)

	answer := part1(input)
	want := 7

	if answer != want {
		t.Errorf("got %d, wanted %d", answer, want)
	}
}

func TestExamplePart2(t *testing.T) {
	input := []byte(`..##.......
#...#...#..
.#....#..#.
..#.#...#.#
.#...##..#.
..#.##.....
.#.#.#....#
.#........#
#.##...#...
#...##....#
.#..#...#.#`)

	answer := part2(input)
	want := 336

	if answer != want {
		t.Errorf("got %d, wanted %d", answer, want)
	}
}
