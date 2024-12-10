package main

import (
	"log"
	"testing"

	"github.com/richardbertozzo/adventofcode-2024/file"
)

func Test_countXmasWords(t *testing.T) {
	content, err := file.ReadFileContent("/input_example")
	if err != nil {
		log.Fatal(err)
	}
	input1 := convertInputToSlice(string(content))

	content2, err := file.ReadFileContent("/input")
	if err != nil {
		log.Fatal(err)
	}
	input2 := convertInputToSlice(string(content2))

	type args struct {
		matrix [][]rune
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "example input - https://adventofcode.com/2024/day/4",
			args: args{
				matrix: input1,
			},
			want: 18,
		},
		{
			name: "my puzzle input - https://adventofcode.com/2024/day/4",
			args: args{
				matrix: input2,
			},
			want: 2397,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := countXmasWords(tt.args.matrix); got != tt.want {
				t.Errorf("countXmasWords() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_countMasInXWords(t *testing.T) {
	content, err := file.ReadFileContent("/input_example_2")

	if err != nil {
		log.Fatal(err)
	}
	input1 := convertInputToSlice(string(content))

	content2, err := file.ReadFileContent("/input")
	if err != nil {
		log.Fatal(err)
	}
	input2 := convertInputToSlice(string(content2))

	type args struct {
		matrix [][]rune
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "example input - https://adventofcode.com/2024/day/4",
			args: args{
				matrix: input1,
			},
			want: 9,
		},
		{
			name: "my puzzle input - https://adventofcode.com/2024/day/4",
			args: args{
				matrix: input2,
			},
			want: 1824,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := countMasInXWords(tt.args.matrix); got != tt.want {
				t.Errorf("countMasInXWords() = %v, want %v", got, tt.want)
			}
		})
	}
}
