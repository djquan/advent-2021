package day2

import (
	"strconv"
	"strings"
)

func part1(input []string) int {
	horizontal, vertical := 0, 0

	for _, s := range input {
		split := strings.Split(s, " ")
		n, err := strconv.Atoi(split[1])

		if err != nil {
			panic("oh no")
		}

		switch split[0] {
		case "forward":
			horizontal += n
		case "down":
			vertical += n
		case "up":
			vertical -= n
		}

	}

	return horizontal * vertical
}

func part2(input []string) int {
	horizontal, vertical, aim := 0, 0, 0

	for _, s := range input {
		split := strings.Split(s, " ")
		n, err := strconv.Atoi(split[1])

		if err != nil {
			panic("oh no")
		}

		switch split[0] {
		case "forward":
			horizontal += n
			vertical += aim * n
		case "down":
			aim += n
		case "up":
			aim -= n
		}

	}

	return horizontal * vertical
}
