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
	oxygen := part2Helper(input, func(a, b rune) bool {
		return a == b
	})

	c02 := part2Helper(input, func(a, b rune) bool {
		return a != b
	})

	return oxygen * c02
}

type filterFn func(a, b rune) bool

func part2Helper(input []string, filter filterFn) int64 {
	candidates := input
	length := len(input[0])

	for i := 0; i < length; i++ {
		if len(candidates) == 1 {
			break
		}

		byteCount := 0
		for _, s := range candidates {
			if s[i] == '1' {
				byteCount++
			}
		}

		var mostCommon rune
		if byteCount > len(candidates)/2 || byteCount == len(candidates)/2 && len(candidates)%2 == 0 {
			mostCommon = '1'
		} else {
			mostCommon = '0'
		}

		newCandidates := make([]string, 0)
		for _, s := range candidates {
			if filter(rune(s[i]), mostCommon) {
				newCandidates = append(newCandidates, s)
			}
		}

		candidates = newCandidates
	}

	result, err := strconv.ParseInt(string(candidates[0]), 2, 64)

	if err != nil {
		panic("OH NO")
	}

	return result
}
