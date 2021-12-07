package day7

import "math"

func part1(input []int) int {
	count := make(map[int]int)

	var min, max int
	for i, x := range input {
		if i == 0 {
			min, max = x, x
		}

		if min > x {
			min = x
		}

		if max < x {
			max = x
		}

		count[x]++
	}

	result := math.MaxInt

	for i := min; i <= max; i++ {
		tmp := 0
		for k, v := range count {
			distance := k - i
			if distance < 0 {
				distance *= -1
			}

			tmp += distance * v
		}

		if tmp < result {
			result = tmp
		}
	}

	return result
}

func part2(input []int) int {
	count := make(map[int]int)

	var min, max int
	for i, x := range input {
		if i == 0 {
			min, max = x, x
		}

		if min > x {
			min = x
		}

		if max < x {
			max = x
		}

		count[x]++
	}

	result := math.MaxInt

	for i := min; i <= max; i++ {
		tmp := 0
		for k, v := range count {
			distance := k - i
			if distance < 0 {
				distance *= -1
			}

			for i := 1; i <= distance; i++ {
				tmp += i * v
			}
		}

		if tmp < result {
			result = tmp
		}
	}

	return result
}
