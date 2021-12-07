package day6

import (
	"advent2021"
	"testing"
)

func Test_part1(t *testing.T) {
	type args struct {
		input []int
		days  int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "One input",
			args: args{
				input: []int{3},
				days:  5,
			},
			want: 2,
		},
		{
			name: "Simple input",
			args: args{
				input: simpleInput(),
				days:  18,
			},
			want: 26,
		},
		{
			name: "Simple input, more days",
			args: args{
				input: simpleInput(),
				days:  80,
			},
			want: 5934,
		},
		{
			name: "Larger Input",
			args: args{
				input: largerInput(),
				days:  80,
			},
			want: 356190,
		},
		{
			name: "Simple input, many more days",
			args: args{
				input: simpleInput(),
				days:  256,
			},
			want: 26984457539,
		},
		{
			name: "Larger Input, many more days",
			args: args{
				input: largerInput(),
				days:  256,
			},
			want: 1617359101538,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := solution(tt.args.input, tt.args.days); got != tt.want {
				t.Errorf("solution() = %v, want %v", got, tt.want)
			}
		})
	}
}

func simpleInput() []int {
	return []int{3, 4, 3, 1, 2}
}

func largerInput() []int {
	return advent2021.ReadIntLine("input.txt")
}
