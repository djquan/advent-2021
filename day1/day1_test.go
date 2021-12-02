package day1

import (
	"bufio"
	"os"
	"strconv"
	"testing"
)

func Test_part1(t *testing.T) {
	type args struct {
		depths []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "Simple test",
			args: args{
				depths: []int{
					199,
					200,
					208,
					210,
					200,
					207,
					240,
					269,
					260,
					263,
				},
			},
			want: 7,
		},
		{
			name: "Empty",
			args: args{
				depths: []int{},
			},
			want: 0,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := part1(tt.args.depths); got != tt.want {
				t.Errorf("part1() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestWithLargerInput(t *testing.T) {
	file, err := os.Open("input.txt")
	if err != nil {
		t.Fatalf("Oh no, the file failed to open: %v", err)
	}
	defer file.Close()

	input := make([]int, 0)
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		parsed, err := strconv.Atoi(scanner.Text())

		if err != nil {
			t.Fatalf("Oh on, the file contained non integers: %v", err)
		}

		input = append(input, parsed)
	}

	result := part1(input)
	if result != 1696 {
		t.Fatalf("Expected 1696, got %v", result)
	}

	result = part2(input)
	if result != 1696 {
		t.Fatalf("Expected 1696, got %v", result)
	}
}

func Test_part2(t *testing.T) {
	type args struct {
		depths []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "Simple",
			args: args{
				depths: []int{
					199,
					200,
					208,
					210,
					200,
					207,
					240,
					269,
					260,
					263,
				},
			},
			want: 5,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := part2(tt.args.depths); got != tt.want {
				t.Errorf("part2() = %v, want %v", got, tt.want)
			}
		})
	}
}
