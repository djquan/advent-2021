package day2

import (
	"bufio"
	"os"
	"testing"
)

func Test_part1(t *testing.T) {
	type args struct {
		input []string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "Simple",
			args: args{
				input: simpleInput(),
			},
			want: 150,
		},
		{
			name: "Larger",
			args: args{
				input: largerInput(),
			},
			want: 2070300,
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
		want int
	}{
		{
			name: "Simple",
			args: args{
				input: simpleInput(),
			},
			want: 900,
		},
		{
			name: "Larger",
			args: args{
				input: largerInput(),
			},
			want: 2078985210,
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
		"forward 5",
		"down 5",
		"forward 8",
		"up 3",
		"down 8",
		"forward 2",
	}
}

func largerInput() []string {
	file, err := os.Open("input.txt")
	if err != nil {
		panic("oh no")
	}
	defer file.Close()

	input := make([]string, 0)
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		input = append(input, scanner.Text())
	}

	return input
}
