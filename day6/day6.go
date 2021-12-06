package day6

func solution(input []int, days int) int {
	fishToDays := make(map[int]int)

	for _, x := range input {
		fishToDays[x]++
	}

	for i := 0; i < days; i++ {
		newFishToDays := make(map[int]int)
		for k, v := range fishToDays {
			if k == 0 {
				newFishToDays[8] = v
				newFishToDays[6] += v
				continue
			}

			newFishToDays[k-1] += v
		}

		fishToDays = newFishToDays
	}

	result := 0
	for _, v := range fishToDays {
		result += v
	}
	return result
}
