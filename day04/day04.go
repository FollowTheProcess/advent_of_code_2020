/*
--- Day 4: Passport Processing ---

You arrive at the airport only to realize that you grabbed your North Pole Credentials instead of your passport. While these documents are extremely similar, North Pole Credentials aren't issued by a country and therefore aren't actually valid documentation for travel in most of the world.

It seems like you're not the only one having problems, though; a very long line has formed for the automatic passport scanners, and the delay could upset your travel itinerary.

Due to some questionable network security, you realize you might be able to solve both of these problems at the same time.

The automatic passport scanners are slow because they're having trouble detecting which passports have all required fields. The expected fields are as follows:

byr (Birth Year)
iyr (Issue Year)
eyr (Expiration Year)
hgt (Height)
hcl (Hair Color)
ecl (Eye Color)
pid (Passport ID)
cid (Country ID)

Passport data is validated in batch files (your puzzle input). Each passport is represented as a sequence of key:value pairs separated by spaces or newlines. Passports are separated by blank lines.

Here is an example batch file containing four passports:

ecl:gry pid:860033327 eyr:2020 hcl:#fffffd
byr:1937 iyr:2017 cid:147 hgt:183cm

iyr:2013 ecl:amb cid:350 eyr:2023 pid:028048884
hcl:#cfa07d byr:1929

hcl:#ae17e1 iyr:2013
eyr:2024
ecl:brn pid:760753108 byr:1931
hgt:179cm

hcl:#cfa07d eyr:2025 pid:166559648
iyr:2011 ecl:brn hgt:59in
The first passport is valid - all eight fields are present. The second passport is invalid - it is missing hgt (the Height field).

The third passport is interesting; the only missing field is cid, so it looks like data from North Pole Credentials, not a passport at all! Surely, nobody would mind if you made the system temporarily ignore missing cid fields. Treat this "passport" as valid.

The fourth passport is missing two fields, cid and byr. Missing cid is fine, but missing any other field is not, so this passport is invalid.

According to the above rules, your improved system would report 2 valid passports.

Count the number of valid passports - those that have all required fields. Treat cid as optional. In your batch file, how many passports are valid?


--- Part Two ---

The line is moving more quickly now, but you overhear airport security talking about how passports with invalid data are getting through. Better add some data validation, quick!

You can continue to ignore the cid field, but each other field has strict rules about what values are valid for automatic validation:

byr (Birth Year) - four digits; at least 1920 and at most 2002.
iyr (Issue Year) - four digits; at least 2010 and at most 2020.
eyr (Expiration Year) - four digits; at least 2020 and at most 2030.
hgt (Height) - a number followed by either cm or in:
If cm, the number must be at least 150 and at most 193.
If in, the number must be at least 59 and at most 76.
hcl (Hair Color) - a # followed by exactly six characters 0-9 or a-f.
ecl (Eye Color) - exactly one of: amb blu brn gry grn hzl oth.
pid (Passport ID) - a nine-digit number, including leading zeroes.
cid (Country ID) - ignored, missing or not.

Your job is to count the passports where all required fields are both present and valid according to the above rules. Here are some example values:

byr valid:   2002
byr invalid: 2003

hgt valid:   60in
hgt valid:   190cm
hgt invalid: 190in
hgt invalid: 190

hcl valid:   #123abc
hcl invalid: #123abz
hcl invalid: 123abc

ecl valid:   brn
ecl invalid: wat

pid valid:   000000001
pid invalid: 0123456789
Here are some invalid passports:

eyr:1972 cid:100
hcl:#18171d ecl:amb hgt:170 pid:186cm iyr:2018 byr:1926

iyr:2019
hcl:#602927 eyr:1967 hgt:170cm
ecl:grn pid:012533040 byr:1946

hcl:dab227 iyr:2012
ecl:brn hgt:182cm pid:021572410 eyr:2020 byr:1992 cid:277

hgt:59cm ecl:zzz
eyr:2038 hcl:74454a iyr:2023
pid:3556412378 byr:2007
Here are some valid passports:

pid:087499704 hgt:74in ecl:grn iyr:2012 eyr:2030 byr:1980
hcl:#623a2f

eyr:2029 ecl:blu cid:129 byr:1989
iyr:2014 pid:896056539 hcl:#a97842 hgt:165cm

hcl:#888785
hgt:164cm byr:2001 iyr:2015 cid:88
pid:545766238 ecl:hzl
eyr:2022

iyr:2010 hgt:158cm hcl:#b6652a ecl:blu byr:1944 eyr:2021 pid:093154719
Count the number of valid passports - those that have all required fields and valid values. Continue to treat cid as optional. In your batch file, how many passports are valid?

*/

package main

import (
	"fmt"
	"os"
	"path/filepath"
	"regexp"
	"strconv"
	"strings"

	"github.com/FollowTheProcess/advent_of_code_2020/utils"
)

type Passport struct {
	BirthYear      int
	IssueYear      int
	ExpirationYear int
	Height         int
	HeightUnit     string
	HairColor      string
	EyeColor       string
	PID            string
	CID            string
}

// IsValid determines if the caller is a valid passport
//
// To be valid, it must have all fields present, with the exception
// of CID which is the only one allowed to be missing
func (p *Passport) IsValid() bool {
	ints := [...]int{p.BirthYear, p.IssueYear, p.ExpirationYear, p.Height}
	// We don't care about CID, it's allowed to be missing
	// and HeightUnit is our own creation
	strs := [...]string{p.HairColor, p.EyeColor, p.PID}

	// Missing int would be the zero value: 0
	for _, i := range ints {
		if i == 0 {
			// Missing, so is invalid
			return false
		}
	}

	// Missing string would be the zero value: ""
	for _, s := range strs {
		if s == "" {
			return false
		}
	}

	// Must be valid
	return true
}

// IsValid2 is the same as IsValid but with part 2 contraints
func (p *Passport) IsValid2() bool {
	// This is a superset of IsValid so mightaswell call that first
	if !p.IsValid() {
		return false
	}

	// Extra part 2 rules

	// Birth year must be between 1920 and 2002
	if p.BirthYear < 1920 || p.BirthYear > 2002 {
		return false
	}

	// Issue year must be between 2010 and 2020
	if p.IssueYear < 2010 || p.IssueYear > 2020 {
		return false
	}

	// Expiration year must be between 2020 and 2030
	if p.ExpirationYear < 2020 || p.ExpirationYear > 2030 {
		return false
	}

	// Height must have either 'cm' or 'in'
	if p.HeightUnit == "" {
		return false
	}

	// If 'cm' height must be between 150 and 193
	if p.HeightUnit == "cm" {
		if p.Height < 150 || p.Height > 193 {
			return false
		}
	}

	// If 'in' height must be between 59 and 76
	if p.HeightUnit == "in" {
		if p.Height < 59 || p.Height > 76 {
			return false
		}
	}

	// Hair color needs to be a valid hex code
	// # followed by 6 characters 0-9 or a-f
	re := regexp.MustCompile(`^#(?:[0-9a-fA-F]{3}){1,2}$`)
	if !re.MatchString(p.HairColor) {
		return false
	}

	// Eye color needs to be certain allowed values
	allowedEyeColors := make(map[string]struct{}, 7)
	for _, item := range [...]string{"amb", "blu", "brn", "gry", "grn", "hzl", "oth"} {
		allowedEyeColors[item] = struct{}{}
	}

	if _, ok := allowedEyeColors[p.EyeColor]; !ok {
		return false
	}

	// PID needs to be 9 digits
	if len(p.PID) != 9 {
		return false
	}

	// Must be valid
	return true
}

// Parse parses the input text and returns a Passport
func Parse(text string) (*Passport, error) {
	parts := strings.Fields(text)

	pairs := make(map[string]string)

	for _, part := range parts {
		subs := strings.Split(part, ":")
		if len(subs) != 2 {
			return nil, fmt.Errorf("key value pair had more than 2 items: %v", subs)
		}
		key, value := subs[0], subs[1]
		pairs[key] = value
	}

	// I tried JSON but couldn't get it to work nicely so let's do it manually
	p := &Passport{}

	if byr, ok := pairs["byr"]; ok {
		birth, err := strconv.Atoi(byr)
		if err != nil {
			return nil, err
		}
		p.BirthYear = birth
	}

	if iyr, ok := pairs["iyr"]; ok {
		issue, err := strconv.Atoi(iyr)
		if err != nil {
			return nil, err
		}
		p.IssueYear = issue
	}

	if eyr, ok := pairs["eyr"]; ok {
		expiry, err := strconv.Atoi(eyr)
		if err != nil {
			return nil, err
		}
		p.ExpirationYear = expiry
	}

	if hgt, ok := pairs["hgt"]; ok {
		// Height is something like "150cm" or "50in"
		switch {
		case strings.Contains(hgt, "cm"):
			h := strings.ReplaceAll(hgt, "cm", "")
			hInt, err := strconv.Atoi(h)
			if err != nil {
				return nil, err
			}
			p.Height = hInt
			p.HeightUnit = "cm"

		case strings.Contains(hgt, "in"):
			h := strings.ReplaceAll(hgt, "in", "")
			hInt, err := strconv.Atoi(h)
			if err != nil {
				return nil, err
			}
			p.Height = hInt
			p.HeightUnit = "in"

		default:
			// Sometimes just bare heights
			h, err := strconv.Atoi(hgt)
			if err != nil {
				return nil, err
			}
			p.Height = h
		}
	}

	if hcl, ok := pairs["hcl"]; ok {
		p.HairColor = hcl
	}

	if ecl, ok := pairs["ecl"]; ok {
		p.EyeColor = ecl
	}

	if pid, ok := pairs["pid"]; ok {
		p.PID = pid
	}

	if cid, ok := pairs["cid"]; ok {
		p.CID = cid
	}

	return p, nil
}

func main() {
	if err := run(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func run() error {
	input, err := os.ReadFile(filepath.Join(utils.Root(), "day04", "day04.txt"))
	if err != nil {
		return err
	}

	items := strings.Split(string(input), "\n\n")

	var passports []*Passport
	for _, item := range items {
		p, err := Parse(item)
		if err != nil {
			return err
		}
		passports = append(passports, p)
	}

	// Count up the valid ones
	valid := 0
	for _, passport := range passports {
		if passport.IsValid() {
			valid++
		}
	}

	// Part 2
	validPart2 := 0
	for _, passport := range passports {
		if passport.IsValid2() {
			validPart2++
		}
	}

	fmt.Printf("Part 1: %d\n", valid)
	fmt.Printf("Part 2: %d\n", validPart2)
	return nil
}
