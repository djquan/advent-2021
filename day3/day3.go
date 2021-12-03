package day3

import "strconv"

func part1(input []string) int64 {
	length := len(input[0])
	byteCount := make([]int, length)

	for _, s := range input {
		for i := 0; i < length; i++ {
			if s[i] == '1' {
				byteCount[i]++
			}
		}
	}

	gamma, eps := make([]rune, length), make([]rune, length)

	for i, count := range byteCount {
		if count > len(input)/2 {
			gamma[i] = '1'
			eps[i] = '0'
		} else {
			gamma[i] = '0'
			eps[i] = '1'
		}
	}

	gammaInt, err := strconv.ParseInt(string(gamma), 2, 64)

	if err != nil {
		panic("OH NO")
	}

	epsInt, err := strconv.ParseInt(string(eps), 2, 64)

	return gammaInt * epsInt
}

func part2(input []string) int64 {
	length := len(input[0])

	candidatesForOxygen := input

	for i := 0; i < length; i++ {
		byteCount := 0
		for _, s := range candidatesForOxygen {
			if s[i] == '1' {
				byteCount++
			}
		}

		var mostCommon rune

		if byteCount > len(candidatesForOxygen)/2 || byteCount == len(candidatesForOxygen)/2 && len(candidatesForOxygen)%2 == 0 {
			mostCommon = '1'
		} else {
			mostCommon = '0'
		}

		newCandidatesForOxygen := make([]string, 0)
		for _, s := range candidatesForOxygen {
			if rune(s[i]) == mostCommon {
				newCandidatesForOxygen = append(newCandidatesForOxygen, s)
			}
		}

		if len(candidatesForOxygen) == 1 {
			break
		}
		candidatesForOxygen = newCandidatesForOxygen
	}

	candidatesForC02 := input
	for i := 0; i < length; i++ {
		byteCount := 0
		for _, s := range candidatesForC02 {
			if s[i] == '1' {
				byteCount++
			}
		}

		var mostCommon rune

		if byteCount > len(candidatesForC02)/2 || byteCount == len(candidatesForC02)/2 && len(candidatesForC02)%2 == 0 {
			mostCommon = '1'
		} else {
			mostCommon = '0'
		}

		newCandidatesForC02 := make([]string, 0)
		for _, s := range candidatesForC02 {
			if rune(s[i]) != mostCommon {
				newCandidatesForC02 = append(newCandidatesForC02, s)
			}
		}

		if len(candidatesForC02) == 1 {
			break
		}
		candidatesForC02 = newCandidatesForC02
	}

	oxygen, err := strconv.ParseInt(string(candidatesForOxygen[0]), 2, 64)

	if err != nil {
		panic("OH NO")
	}

	c02, err := strconv.ParseInt(string(candidatesForC02[0]), 2, 64)

	return oxygen * c02
}
