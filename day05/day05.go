/*
--- Day 5: Binary Boarding ---

You board your plane only to discover a new problem: you dropped your boarding pass! You aren't sure which seat is yours, and all of the flight attendants are busy with the flood of people that suddenly made it through passport control.

You write a quick program to use your phone's camera to scan all of the nearby boarding passes (your puzzle input); perhaps you can find your seat through process of elimination.

Instead of zones or groups, this airline uses binary space partitioning to seat people. A seat might be specified like FBFBBFFRLR, where F means "front", B means "back", L means "left", and R means "right".

The first 7 characters will either be F or B; these specify exactly one of the 128 rows on the plane (numbered 0 through 127). Each letter tells you which half of a region the given seat is in. Start with the whole list of rows; the first letter indicates whether the seat is in the front (0 through 63) or the back (64 through 127). The next letter indicates which half of that region the seat is in, and so on until you're left with exactly one row.

For example, consider just the first seven characters of FBFBBFFRLR:

Start by considering the whole range, rows 0 through 127.
F means to take the lower half, keeping rows 0 through 63.
B means to take the upper half, keeping rows 32 through 63.
F means to take the lower half, keeping rows 32 through 47.
B means to take the upper half, keeping rows 40 through 47.
B keeps rows 44 through 47.
F keeps rows 44 through 45.

The final F keeps the lower of the two, row 44.
The last three characters will be either L or R; these specify exactly one of the 8 columns of seats on the plane (numbered 0 through 7). The same process as above proceeds again, this time with only three steps. L means to keep the lower half, while R means to keep the upper half.

For example, consider just the last 3 characters of FBFBBFFRLR:

Start by considering the whole range, columns 0 through 7.
R means to take the upper half, keeping columns 4 through 7.
L means to take the lower half, keeping columns 4 through 5.
The final R keeps the upper of the two, column 5.
So, decoding FBFBBFFRLR reveals that it is the seat at row 44, column 5.

Every seat also has a unique seat ID: multiply the row by 8, then add the column. In this example, the seat has ID 44 * 8 + 5 = 357.

Here are some other boarding passes:

BFFFBBFRRR: row 70, column 7, seat ID 567.
FFFBBBFRRR: row 14, column 7, seat ID 119.
BBFFBBFRLL: row 102, column 4, seat ID 820.

As a sanity check, look through your list of boarding passes. What is the highest seat ID on a boarding pass?

--- Part Two ---

Ding! The "fasten seat belt" signs have turned on. Time to find your seat.

It's a completely full flight, so your seat should be the only missing boarding pass in your list. However, there's a catch: some of the seats at the very front and back of the plane don't exist on this aircraft, so they'll be missing from your list as well.

Your seat wasn't at the very front or back, though; the seats with IDs +1 and -1 from yours will be in your list.

What is the ID of your seat?

*/
package main

import (
	"fmt"
	"os"
	"path/filepath"
	"sort"
	"strings"

	"github.com/FollowTheProcess/advent_of_code_2020/utils"
)

const (
	maxRow = 127
	maxCol = 7
)

func main() {
	if err := run(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func run() error {
	data, err := os.ReadFile(filepath.Join(utils.Root(), "day05", "day05.txt"))
	if err != nil {
		return err
	}

	max := 0
	seats := []int{}

	for _, line := range strings.Split(string(data), "\n") {
		if len(line) == 0 {
			continue
		}

		row, col, err := findSeat(line)
		if err != nil {
			return err
		}

		seatID := row*8 + col
		seats = append(seats, seatID)

		if seatID > max {
			max = seatID
		}
	}

	fmt.Printf("Part 1: %d\n", max)

	// Find the missing seat by sorting and finding the gap
	sort.Ints(seats)
	prev := seats[0]
	for i := 1; i < len(seats); i++ {
		if seats[i] != prev+1 {
			// We've found the gap!
			fmt.Printf("Part 2: %d\n", prev+1)
			break
		}
		prev = seats[i]
	}

	return nil
}

func findSeat(code string) (row, col int, err error) {
	// First, find the row
	start := 0
	end := maxRow

	// First 7 chars of the code refer to row
	for _, char := range code[:7] {
		middle := (start + end + 1) / 2

		switch char {
		case 'B':
			// Back half i.e. [middle:]
			start = middle
		case 'F':
			// Front half i.e. [:middle]
			end = middle - 1
		default:
			return 0, 0, fmt.Errorf("expected 'F' or 'B', got %s", string(char))
		}
	}

	row = start

	start = 0
	end = maxCol
	// Last 3 characters refer to column
	for _, char := range code[7:] {
		middle := (start + end + 1) / 2

		switch char {
		case 'R':
			// Right half i.e. [middle:]
			start = middle
		case 'L':
			// Left half, i.e. [:middle]
			end = middle - 1
		default:
			return 0, 0, fmt.Errorf("expected 'L' or 'R', got %s", string(char))
		}
	}

	col = start

	return row, col, nil
}
