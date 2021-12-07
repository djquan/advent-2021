package day7

import (
	"advent2021"
	"testing"
)

func Test_part1(t *testing.T) {
	type args struct {
		input []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "Simple input",
			args: args{
				input: simpleInput(),
			},
			want: 37,
		},
		{
			name: "Larger input",
			args: args{
				input: largerInput(),
			},
			want: 352997,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := part1(tt.args.input); got != tt.want {
				t.Errorf("part1() = %v, want %v", got, tt.want)
			}
		})
	}
}

func simpleInput() []int {
	return []int{16, 1, 2, 0, 4, 2, 7, 1, 2, 14}
}

func largerInput() []int {
	return advent2021.ReadIntLine("input.txt")
}

func Test_part2(t *testing.T) {
	type args struct {
		input []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "Simple input",
			args: args{
				input: simpleInput(),
			},
			want: 168,
		},
		{
			name: "Larger input",
			args: args{
				input: largerInput(),
			},
			want: 101571302,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := part2(tt.args.input); got != tt.want {
				t.Errorf("part2() = %v, want %v", got, tt.want)
			}
		})
	}
}
