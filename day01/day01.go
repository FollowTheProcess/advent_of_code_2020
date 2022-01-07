/*
--- Day 1: Report Repair ---

After saving Christmas five years in a row, you've decided to take a vacation at a nice resort on a tropical island. Surely, Christmas will go on without you.

The tropical island has its own currency and is entirely cash-only. The gold coins used there have a little picture of a starfish; the locals just call them stars. None of the currency exchanges seem to have heard of them, but somehow, you'll need to find fifty of these coins by the time you arrive so you can pay the deposit on your room.

To save your vacation, you need to get all fifty stars by December 25th.

Collect stars by solving puzzles. Two puzzles will be made available on each day in the Advent calendar; the second puzzle is unlocked when you complete the first. Each puzzle grants one star. Good luck!

Before you leave, the Elves in accounting just need you to fix your expense report (your puzzle input); apparently, something isn't quite adding up.

Specifically, they need you to find the two entries that sum to 2020 and then multiply those two numbers together.

For example, suppose your expense report contained the following:

1721
979
366
299
675
1456

In this list, the two entries that sum to 2020 are 1721 and 299. Multiplying them together produces 1721 * 299 = 514579, so the correct answer is 514579.

Of course, your expense report is much larger. Find the two entries that sum to 2020; what do you get if you multiply them together?

--- Part Two ---

The Elves in accounting are thankful for your help; one of them even offers you a starfish coin they had left over from a past vacation. They offer you a second one if you can find three numbers in your expense report that meet the same criteria.

Using the above example again, the three entries that sum to 2020 are 979, 366, and 675. Multiplying them together produces the answer, 241861950.

In your expense report, what is the product of the three entries that sum to 2020?
*/
package main

import (
	"fmt"
	"os"
	"path/filepath"
	"strconv"
	"strings"

	"github.com/FollowTheProcess/advent_of_code_2020/hashset"
	"github.com/FollowTheProcess/advent_of_code_2020/utils"
)

func main() {
	if err := run(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func run() error {
	inputFile := filepath.Join(utils.Root(), "day01", "day01.txt")
	input, err := os.ReadFile(inputFile)
	if err != nil {
		return err
	}

	inStr := string(input)
	entries := strings.Split(inStr, "\n")

	intEntries := []int{}
	for _, i := range entries {
		entry, _ := strconv.Atoi(i)
		intEntries = append(intEntries, entry)
	}

	fmt.Printf("Part 1: %d\n\n", part1(intEntries))
	fmt.Printf("Part 2: %d\n", part2(intEntries))
	return nil
}

func part1(entries []int) int {
	original := hashset.IntHashSet{}
	for _, entry := range entries {
		original.Add(entry)
	}
	// What needs to be added to the original to make 2020
	needs := hashset.IntHashSet{}
	for _, entry := range entries {
		needs.Add(2020 - entry)
	}

	// See if any of these diffs are in the original set
	matches := []int{}
	for need := range needs {
		if original.Has(need) {
			matches = append(matches, need)
		}
	}

	// Return the product of the things that add up to 2020
	// according to the problem there should only be two
	if len(matches) > 2 {
		panic("Matches should only contain 2 elements")
	}

	return matches[0] * matches[1]
}

func part2(entries []int) int {
	original := hashset.IntHashSet{}
	for _, entry := range entries {
		original.Add(entry)
	}

	needs := make(map[int][]int)
	for _, i := range entries {
		for _, j := range entries {
			if i != j {
				needs[2020-i-j] = []int{i, j}
			}
		}
	}

	for _, entry := range entries {
		if n, ok := needs[entry]; ok {
			j, k := n[0], n[1]
			return entry * j * k
		}
	}

	return 0
}
