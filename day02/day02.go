/*
--- Day 2: Password Philosophy ---

Your flight departs in a few days from the coastal airport; the easiest way down to the coast from here is via toboggan.

The shopkeeper at the North Pole Toboggan Rental Shop is having a bad day. "Something's wrong with our computers; we can't log in!" You ask if you can take a look.

Their password database seems to be a little corrupted: some of the passwords wouldn't have been allowed by the Official Toboggan Corporate Policy that was in effect when they were chosen.

To try to debug the problem, they have created a list (your puzzle input) of passwords (according to the corrupted database) and the corporate policy when that password was set.

For example, suppose you have the following list:

1-3 a: abcde
1-3 b: cdefg
2-9 c: ccccccccc
Each line gives the password policy and then the password. The password policy indicates the lowest and highest number of times a given letter must appear for the password to be valid. For example, 1-3 a means that the password must contain a at least 1 time and at most 3 times.

In the above example, 2 passwords are valid. The middle password, cdefg, is not; it contains no instances of b, but needs at least 1. The first and third passwords are valid: they contain one a or nine c, both within the limits of their respective policies.

How many passwords are valid according to their policies?

--- Part Two ---

While it appears you validated the passwords correctly, they don't seem to be what the Official Toboggan Corporate Authentication System is expecting.

The shopkeeper suddenly realizes that he just accidentally explained the password policy rules from his old job at the sled rental place down the street! The Official Toboggan Corporate Policy actually works a little differently.

Each policy actually describes two positions in the password, where 1 means the first character, 2 means the second character, and so on. (Be careful; Toboggan Corporate Policies have no concept of "index zero"!) Exactly one of these positions must contain the given letter. Other occurrences of the letter are irrelevant for the purposes of policy enforcement.

Given the same example list from above:

1-3 a: abcde is valid: position 1 contains a and position 3 does not.
1-3 b: cdefg is invalid: neither position 1 nor position 3 contains b.
2-9 c: ccccccccc is invalid: both position 2 and position 9 contain c.
How many passwords are valid according to the new interpretation of the policies?

*/
package main

import (
	"fmt"
	"os"
	"path/filepath"
	"runtime"
	"strconv"
	"strings"
)

type Password struct {
	Text   string
	Letter string
	Min    int
	Max    int
}

// Parse takes the raw line e.g. '1-3 a: abcde' and returns a Password
func Parse(raw string) (*Password, error) {
	parts := strings.Split(raw, " ") // []string{"1-3", "a:", "abcde"}

	// Should always be 3 parts
	if len(parts) != 3 {
		return nil, fmt.Errorf("malformed password string: %s", raw)
	}

	minMax := strings.Split(parts[0], "-")
	if len(minMax) != 2 {
		return nil, fmt.Errorf("expected 2 parts in constraint: %s", parts[0])
	}

	min, err := strconv.Atoi(minMax[0])
	if err != nil {
		return nil, fmt.Errorf("min part of constraint not a valid integer: %v", minMax[0])
	}

	max, err := strconv.Atoi(minMax[1])
	if err != nil {
		return nil, fmt.Errorf("max part of constraint not a valid integer: %v", minMax[1])
	}

	letter := strings.Replace(parts[1], ":", "", 1)

	p := &Password{
		Text:   parts[2],
		Letter: letter,
		Min:    min,
		Max:    max,
	}

	return p, nil
}

// IsValid checks whether the password is valid according to it's policy
func (p *Password) IsValid() bool {
	count := strings.Count(p.Text, p.Letter)

	if count >= p.Min && count <= p.Max {
		return true
	}

	return false
}

// IsValidPart2 checks whether the password meets the part 2 policy
func (p *Password) IsValidPart2() bool {
	// Compensate for no zero index, 1 actually means 0 etc
	minIndex, maxIndex := p.Min-1, p.Max-1

	// Only 1 of these positions can contain letter
	// If both, then false
	if string(p.Text[minIndex]) == p.Letter && string(p.Text[maxIndex]) == p.Letter {
		return false
	}

	// If neither then false
	if string(p.Text[minIndex]) != p.Letter && string(p.Text[maxIndex]) != p.Letter {
		return false
	}

	// Anything else is true
	return true
}

func main() {
	if err := run(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func run() error {
	inputFile := filepath.Join(here(), "day02.txt")
	raw, err := os.ReadFile(inputFile)
	if err != nil {
		return err
	}

	lines := strings.FieldsFunc(string(raw), func(r rune) bool { return r == '\n' })

	var passwords []*Password
	for _, line := range lines {
		p, err := Parse(line)
		if err != nil {
			return err
		}
		passwords = append(passwords, p)
	}

	// Count the valid ones
	valid := 0
	for _, password := range passwords {
		if password.IsValid() {
			valid++
		}
	}

	validPart2 := 0
	for _, password := range passwords {
		if password.IsValidPart2() {
			validPart2++
		}
	}

	fmt.Printf("Part 1: %d\n\n", valid)
	fmt.Printf("Part 2: %d\n", validPart2)

	return nil
}

func here() string {
	_, file, _, ok := runtime.Caller(0)
	if !ok {
		panic("Couldn't find file!")
	}
	return filepath.Join(file, "../")
}
