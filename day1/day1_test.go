package day1

import (
	"advent2021"
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
				depths: simpleInput(),
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
		{
			name: "Larger input",
			args: args{
				depths: largerInput(),
			},
			want: 1696,
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
				depths: simpleInput(),
			},
			want: 5,
		},
		{
			name: "Larger input",
			args: args{
				depths: largerInput(),
			},
			want: 1737,
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

func simpleInput() []int {
	return []int{
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
	}
}

func largerInput() []int {
	input := advent2021.ReadLines("input.txt")

	result := make([]int, 0)

	for _, s := range input {
		parsed, err := strconv.Atoi(s)
		if err != nil {
			panic("oh no")
		}
		result = append(result, parsed)

	}
	return result
}
