package day4

import (
	"advent2021"
	"strconv"
	"strings"
	"testing"
)

func Test_part1(t *testing.T) {
	type args struct {
		input        [][][]int
		instructions []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "Simple",
			args: args{
				input:        simpleInput(),
				instructions: simpleInstructions(),
			},
			want: 4512,
		},
		{
			name: "Larger input",
			args: args{
				input:        largerInput(),
				instructions: largerInstructions(),
			},
			want: 67716,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := part1(tt.args.input, tt.args.instructions); got != tt.want {
				t.Errorf("part1() = %v, want %v", got, tt.want)
			}
		})
	}
}

func largerInstructions() []int {
	return []int{
		31,
		88,
		35,
		24,
		46,
		48,
		95,
		42,
		18,
		43,
		71,
		32,
		92,
		62,
		97,
		63,
		50,
		2,
		60,
		58,
		74,
		66,
		15,
		87,
		57,
		34,
		14,
		3,
		54,
		93,
		75,
		22,
		45,
		10,
		56,
		12,
		83,
		30,
		8,
		76,
		1,
		78,
		82,
		39,
		98,
		37,
		19,
		26,
		81,
		64,
		55,
		41,
		16,
		4,
		72,
		5,
		52,
		80,
		84,
		67,
		21,
		86,
		23,
		91,
		0,
		68,
		36,
		13,
		44,
		20,
		69,
		40,
		90,
		96,
		27,
		77,
		38,
		49,
		94,
		47,
		9,
		65,
		28,
		59,
		79,
		6,
		29,
		61,
		53,
		11,
		17,
		73,
		99,
		25,
		89,
		51,
		7,
		33,
		85,
		70,
	}
}

func largerInput() [][][]int {
	text := advent2021.ReadLines("input.txt")

	result := make([][][]int, 0)
	board := make([][]int, 5)
	currentRow := 0
	for _, s := range text {
		if s == "" {
			result = append(result, board)
			board = make([][]int, 5)
			currentRow = 0
			continue
		}

		board[currentRow] = make([]int, 5)
		elements := strings.Fields(s)

		for i, element := range elements {
			result, err := strconv.Atoi(element)
			if err != nil {
				panic("oh no")
			}

			board[currentRow][i] = result
		}

		currentRow++
	}

	return result
}

func simpleInstructions() []int {
	return []int{
		7,
		4,
		9,
		5,
		11,
		17,
		23,
		2,
		0,
		14,
		21,
		24,
		10,
		16,
		13,
		6,
		15,
		25,
		12,
		22,
		18,
		20,
		8,
		19,
		3,
		26,
		1,
	}

}

func simpleInput() [][][]int {
	a := [][]int{
		{22, 13, 17, 11, 0},
		{8, 2, 23, 4, 24},
		{21, 9, 14, 16, 7},
		{6, 10, 3, 18, 5},
		{1, 12, 20, 15, 19},
	}

	b := [][]int{
		{3, 15, 0, 2, 22},
		{9, 18, 13, 17, 5},
		{19, 8, 7, 25, 23},
		{20, 11, 10, 24, 4},
		{14, 21, 16, 12, 6},
	}

	c := [][]int{
		{14, 21, 17, 24, 4},
		{10, 16, 15, 9, 19},
		{18, 8, 23, 26, 20},
		{22, 11, 13, 6, 5},
		{2, 0, 12, 3, 7},
	}

	return [][][]int{a, b, c}
}

func Test_part2(t *testing.T) {
	type args struct {
		input        [][][]int
		instructions []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "Simple",
			args: args{
				input:        simpleInput(),
				instructions: simpleInstructions(),
			},
			want: 1924,
		},
		{
			name: "Larger input",
			args: args{
				input:        largerInput(),
				instructions: largerInstructions(),
			},
			want: 1830,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := part2(tt.args.input, tt.args.instructions); got != tt.want {
				t.Errorf("part2() = %v, want %v", got, tt.want)
			}
		})
	}
}
