package day3

import (
	"advent2021"
	"testing"
)

func Test_part1(t *testing.T) {
	type args struct {
		input []string
	}
	tests := []struct {
		name string
		args args
		want int64
	}{
		{
			name: "Simple",
			args: args{
				input: simpleInput(),
			},
			want: 198,
		},
		{
			name: "Larger",
			args: args{
				input: largerInput(),
			},
			want: 2640986,
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

func Test_part2(t *testing.T) {
	type args struct {
		input []string
	}
	tests := []struct {
		name string
		args args
		want int64
	}{
		{
			name: "Simple",
			args: args{
				input: simpleInput(),
			},
			want: 230,
		},
		{
			name: "Larger",
			args: args{
				input: largerInput(),
			},
			want: 6822109,
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

func simpleInput() []string {
	return []string{
		"00100",
		"11110",
		"10110",
		"10111",
		"10101",
		"01111",
		"00111",
		"11100",
		"10000",
		"11001",
		"00010",
		"01010",
	}
}

func largerInput() []string {
	return advent2021.ReadLines("input.txt")
}
