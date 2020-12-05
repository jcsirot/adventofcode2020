package main

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"regexp"
	"strconv"
	"strings"
)

type passport struct {
	Byr string
	Iyr string
	Eyr string
	Hgt string
	Hcl string
	Ecl string
	Pid string
	Cid string
}

func newPassport(line string) passport {
	p := passport{}
	r := regexp.MustCompile("([a-z]+):([a-zA-Z0-9#]+)")
	match := r.FindAllStringSubmatch(line, -1)
	for _, f := range match {
		if f[1] == "byr" {
			p.Byr = f[2]
		} else if f[1] == "iyr" {
			p.Iyr = f[2]
		} else if f[1] == "eyr" {
			p.Eyr = f[2]
		} else if f[1] == "hgt" {
			p.Hgt = f[2]
		} else if f[1] == "hcl" {
			p.Hcl = f[2]
		} else if f[1] == "ecl" {
			p.Ecl = f[2]
		} else if f[1] == "pid" {
			p.Pid = f[2]
		} else if f[1] == "cid" {
			p.Cid = f[2]
		}
	}
	return p
}

func (p passport) isValid() bool {
	return p.Byr != "" && p.Iyr != "" && p.Eyr != "" && p.Hgt != "" && p.Hcl != "" && p.Ecl != "" && p.Pid != ""
}

func (p passport) isValidStrict() bool {
	if !p.isValid() {
		return false
	}

	if !isValidNumber(p.Byr, 1920, 2002) {
		return false
	}

	if !isValidNumber(p.Iyr, 2010, 2020) {
		return false
	}

	if !isValidNumber(p.Eyr, 2020, 2030) {
		return false
	}

	if !((strings.HasSuffix(p.Hgt, "cm") && isValidNumber(strings.TrimSuffix(p.Hgt, "cm"), 150, 193)) ||
		(strings.HasSuffix(p.Hgt, "in") && isValidNumber(strings.TrimSuffix(p.Hgt, "in"), 59, 76))) {
		return false
	}

	if !regexp.MustCompile("^#[0-9a-f]{6}$").MatchString(p.Hcl) {
		return false
	}

	if !(p.Ecl == "amb" || p.Ecl == "blu" || p.Ecl == "brn" || p.Ecl == "gry" || p.Ecl == "grn" || p.Ecl == "hzl" || p.Ecl == "oth") {
		return false
	}

	if !regexp.MustCompile("^[0-9]{9}$").MatchString(p.Pid) {
		return false
	}

	return true
}

func isValidNumber(s string, min, max int) bool {
	n, err := strconv.Atoi(s)
	if err != nil {
		return false
	}
	if n < min || n > max {
		return false
	}
	return true
}

func readPassports(path string) ([]passport, error) {
	buf, err := ioutil.ReadFile(path)
	if err != nil {
		return nil, err
	}
	// Yes it's ugly, but it works!
	batch := strings.Replace(string(buf), "\n", " ", -1)
	batch = strings.Replace(batch, "  ", "\n", -1)

	passports := make([]passport, 0)

	scanner := bufio.NewScanner(strings.NewReader(batch))
	for scanner.Scan() {
		passports = append(passports, newPassport(scanner.Text()))
	}
	return passports, nil
}

func part1(passports []passport) int {
	valid := 0
	for _, p := range passports {
		if p.isValid() {
			valid++
		}
	}
	return valid
}

func part2(passports []passport) int {
	valid := 0
	for _, p := range passports {
		if p.isValidStrict() {
			valid++
		}
	}
	return valid
}

func main() {
	passports, _ := readPassports("day04.txt")
	fmt.Printf("Part 1: %d\n", part1(passports))
	fmt.Printf("Part 2: %d\n", part2(passports))
}
