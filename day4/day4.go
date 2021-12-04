package day4

func part1(input [][][]int, instructions []int) int {
	markedTracking := buildTrackingStructure(input)

	for _, instruction := range instructions {
		for bi, board := range input {
			for ri, row := range board {
				for ci, elem := range row {
					if elem == instruction {
						markedTracking[bi][ri][ci] = true

						if won(markedTracking[bi], ri, ci) {
							return elem * sumOfUntracked(markedTracking[bi], board)
						}
					}
				}
			}

		}
	}

	return 0
}

func part2(input [][][]int, instructions []int) int {
	markedTracking := buildTrackingStructure(input)
	wonBoards := map[int]bool{}

	for _, instruction := range instructions {
		for bi, board := range input {
			if wonBoards[bi] {
				continue
			}

			for ri, row := range board {
				for ci, elem := range row {
					if elem == instruction {
						markedTracking[bi][ri][ci] = true

						if won(markedTracking[bi], ri, ci) {
							wonBoards[bi] = true

							if len(wonBoards) == len(input) {
								return elem * sumOfUntracked(markedTracking[bi], board)
							}
						}
					}
				}
			}

		}
	}

	return 0
}

func sumOfUntracked(tracking [][]bool, board [][]int) int {
	result := 0

	for ri, row := range tracking {
		for ci, elem := range row {
			if !elem {
				result += board[ri][ci]
			}

		}
	}

	return result
}

func won(tracking [][]bool, rowIndex int, columnIndex int) bool {
	markedSameRow, markedSameColumn := 0, 0
	for i := 0; i < 5; i++ {
		if tracking[rowIndex][i] {
			markedSameRow++
		}

		if tracking[i][columnIndex] {
			markedSameColumn++
		}
	}

	return markedSameRow == 5 || markedSameColumn == 5
}

func buildTrackingStructure(input [][][]int) [][][]bool {
	boards := make([][][]bool, len(input))

	for i := range boards {
		boards[i] = make([][]bool, 5)
		for j := range boards[i] {
			boards[i][j] = make([]bool, 5)
		}
	}

	return boards
}
