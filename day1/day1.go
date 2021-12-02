package day1

func part1(depths []int) int {
	if len(depths) == 0 {
		return 0
	}

	previous, depths := depths[0], depths[1:]
	result := 0

	for _, depth := range depths {
		if depth > previous {
			result += 1
		}
		previous = depth
	}
	return result
}

func part2(depths []int) int {
	var slidingDepths []int

	for i, depth := range depths {
		if i >= 2 {
			val := depth + depths[i-1] + depths[i-2]
			slidingDepths = append(slidingDepths, val)
		}
	}

	return part1(slidingDepths)
}
