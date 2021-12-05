package day5

import (
	"advent2021"
	"reflect"
	"strconv"
	"strings"
	"testing"
)

func Test_part1(t *testing.T) {
	type args struct {
		input []lineSegment
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
			want: 5,
		},
		{
			name: "Larger input",
			args: args{
				input: largerInput(),
			},
			want: 7674,
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

func simpleInput() []lineSegment {
	input := []string{
		"0,9 -> 5,9",
		"8,0 -> 0,8",
		"9,4 -> 3,4",
		"2,2 -> 2,1",
		"7,0 -> 7,4",
		"6,4 -> 2,0",
		"0,9 -> 2,9",
		"3,4 -> 1,4",
		"0,0 -> 8,8",
		"5,5 -> 8,2",
	}

	return parseLineSegments(input)
}

func largerInput() []lineSegment {
	input := advent2021.ReadLines("input.txt")
	return parseLineSegments(input)
}

func parseLineSegments(input []string) []lineSegment {
	result := make([]lineSegment, len(input))

	for i, s := range input {
		segments := strings.Split(s, "->")
		var segment lineSegment

		first := strings.Split(segments[0], ",")
		convert, err := strconv.Atoi(strings.TrimSpace(first[0]))
		if err != nil {
			panic("oh no")
		}
		segment.first.x = convert

		convert, err = strconv.Atoi(strings.TrimSpace(first[1]))
		if err != nil {
			panic("oh no")
		}
		segment.first.y = convert

		second := strings.Split(segments[1], ",")
		convert, err = strconv.Atoi(strings.TrimSpace(second[0]))
		if err != nil {
			panic("oh no")
		}
		segment.second.x = convert

		convert, err = strconv.Atoi(strings.TrimSpace(second[1]))
		if err != nil {
			panic("oh no")
		}
		segment.second.y = convert

		result[i] = segment
	}

	return result
}

func Test_part2(t *testing.T) {
	type args struct {
		input []lineSegment
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
			want: 12,
		},
		{
			name: "Larger input",
			args: args{
				input: largerInput(),
			},
			want: 20898,
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

func Test_diagonalCoordinates(t *testing.T) {
	type args struct {
		a coordinate
		b coordinate
	}
	tests := []struct {
		name    string
		args    args
		want    []coordinate
		wantErr bool
	}{
		{
			name: "Simple",
			args: args{
				a: coordinate{
					x: 1,
					y: 1,
				},
				b: coordinate{
					x: 3,
					y: 3,
				},
			},
			want: []coordinate{{
				x: 1,
				y: 1,
			}, {
				x: 2,
				y: 2,
			}, {
				x: 3,
				y: 3,
			}},
		},
		{
			name: "another one",
			args: args{
				a: coordinate{
					x: 7,
					y: 9,
				},
				b: coordinate{
					x: 9,
					y: 7,
				},
			},
			want: []coordinate{{
				x: 7,
				y: 9,
			}, {
				x: 8,
				y: 8,
			}, {
				x: 9,
				y: 7,
			}},
		},
		{
			name: "Reversed one",
			args: args{
				a: coordinate{
					x: 9,
					y: 7,
				},
				b: coordinate{
					x: 7,
					y: 9,
				},
			},
			want: []coordinate{{
				x: 7,
				y: 9,
			}, {
				x: 8,
				y: 8,
			}, {
				x: 9,
				y: 7,
			}},
		},
		{
			name: "No diagonal",
			args: args{
				a: coordinate{
					x: 9,
					y: 7,
				},
				b: coordinate{
					x: 7,
					y: 10,
				},
			},
			want: []coordinate{},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := diagonalCoordinates(tt.args.a, tt.args.b)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("diagonalCoordinates() got = %v, want %v", got, tt.want)
			}
		})
	}
}
