package day5

type lineSegment struct {
	first  coordinate
	second coordinate
}

type coordinate struct {
	x, y int
}

func part1(input []lineSegment) int {
	seen := map[coordinate]int{}
	result := 0

	for _, segment := range input {
		var min, max int
		if segment.first.y == segment.second.y {
			if segment.first.x < segment.second.x {
				min = segment.first.x
				max = segment.second.x
			} else {
				min = segment.second.x
				max = segment.first.x
			}

			for i := min; i <= max; i++ {
				coordinate := coordinate{
					x: i,
					y: segment.first.y,
				}

				seen[coordinate]++
				if seen[coordinate] == 2 {
					result++
				}
			}
		}

		if segment.first.x == segment.second.x {
			if segment.first.y < segment.second.y {
				min = segment.first.y
				max = segment.second.y
			} else {
				min = segment.second.y
				max = segment.first.y
			}

			for i := min; i <= max; i++ {
				coordinate := coordinate{
					x: segment.first.x,
					y: i,
				}
				seen[coordinate]++
				if seen[coordinate] == 2 {
					result++
				}
			}
		}
	}

	return result
}

func part2(input []lineSegment) int {
	seen := map[coordinate]int{}
	result := 0

	for _, segment := range input {
		var min, max int
		if segment.first.y == segment.second.y {
			if segment.first.x < segment.second.x {
				min = segment.first.x
				max = segment.second.x
			} else {
				min = segment.second.x
				max = segment.first.x
			}

			for i := min; i <= max; i++ {
				coordinate := coordinate{
					x: i,
					y: segment.first.y,
				}

				seen[coordinate]++
				if seen[coordinate] == 2 {
					result++
				}
			}
			continue
		}

		if segment.first.x == segment.second.x {
			if segment.first.y < segment.second.y {
				min = segment.first.y
				max = segment.second.y
			} else {
				min = segment.second.y
				max = segment.first.y
			}

			for i := min; i <= max; i++ {
				coordinate := coordinate{
					x: segment.first.x,
					y: i,
				}
				seen[coordinate]++
				if seen[coordinate] == 2 {
					result++
				}
			}
			continue
		}

		diagonals := diagonalCoordinates(segment.first, segment.second)

		for _, coordinate := range diagonals {
			seen[coordinate]++
			if seen[coordinate] == 2 {
				result++
			}
		}

	}

	return result
}

func diagonalCoordinates(a, b coordinate) []coordinate {
	var smallestX, largestX coordinate
	if a.x < b.x {
		smallestX = a
		largestX = b
	} else {
		smallestX = b
		largestX = a
	}

	var tmp []coordinate
	for i := 0; i <= largestX.x-smallestX.x; i++ {
		x := smallestX.x + i

		var y int
		if smallestX.y < largestX.y {
			y = smallestX.y + i
		} else {
			y = smallestX.y - i
		}

		tmp = append(tmp, coordinate{x: x, y: y})

		if x == largestX.x && y == largestX.y {
			return tmp
		}
	}

	return []coordinate{}
}
