package advent2021

import (
	"bufio"
	"os"
	"strconv"
	"strings"
)

func ReadLines(filename string) []string {
	file, err := os.Open(filename)
	if err != nil {
		panic("oh no")
	}
	defer file.Close()

	input := make([]string, 0)
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		input = append(input, scanner.Text())
	}

	return input
}

func ReadIntLine(filename string) []int {
	line := ReadLines(filename)[0]

	split := strings.Split(line, ",")

	result := make([]int, len(split))

	for i, s := range split {
		x, err := strconv.Atoi(s)
		if err != nil {
			panic("oh no")
		}

		result[i] = x
	}
	return result
}
