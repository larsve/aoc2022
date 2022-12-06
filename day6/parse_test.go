package main

import (
	"bytes"
	_ "embed"
	"fmt"
	"testing"
)

//go:embed input.txt
var benchInput []byte

func Test_part1Ofs(t *testing.T) {
	tests := []struct {
		input []byte
		want  int
	}{
		{input: nil, want: -1},
		{input: []byte("a"), want: -1},
		{input: []byte("aaaa"), want: -1},
		{input: []byte("mjqjpqmgbljsphdztnvjfqwrcgsmlb"), want: 7},
		{input: []byte("bvwbjplbgvbhsrlpgdmjqwftvncz"), want: 5},
		{input: []byte("nppdvjthqldpwncqszvftbrmjlhg"), want: 6},
		{input: []byte("nznrnfrfntjfmvfwmzdfjlvtqnbhcprsg"), want: 10},
		{input: []byte("zcfzfwzzqfrljwzlrfnpqdbhtmscgvjw"), want: 11},
		{input: benchInput, want: 1198},
	}
	for i, tt := range tests {
		t.Run(fmt.Sprintf("#%d", i), func(t *testing.T) {
			if got := part1Ofs(bytes.NewReader(tt.input)); got != tt.want {
				t.Errorf("ERROR: got: %v, want: %v", got, tt.want)
			}
		})
	}
}

func Benchmark_part1Ofs(b *testing.B) {
	for i := 0; i < b.N; i++ {
		if c := part1Ofs(bytes.NewReader(benchInput)); c != 1198 {
			b.Fatalf("ERROR: wrong offset: %d", c)
		}
	}
}

func Test_part2Ofs(t *testing.T) {
	tests := []struct {
		input []byte
		want  int
	}{
		{input: nil, want: -1},
		{input: []byte("a"), want: -1},
		{input: []byte("aaaaaaaaaaaaaa"), want: -1},
		{input: []byte("mjqjpqmgbljsphdztnvjfqwrcgsmlb"), want: 19},
		{input: []byte("bvwbjplbgvbhsrlpgdmjqwftvncz"), want: 23},
		{input: []byte("nppdvjthqldpwncqszvftbrmjlhg"), want: 23},
		{input: []byte("nznrnfrfntjfmvfwmzdfjlvtqnbhcprsg"), want: 29},
		{input: []byte("zcfzfwzzqfrljwzlrfnpqdbhtmscgvjw"), want: 26},
		{input: benchInput, want: 3120},
	}
	for i, tt := range tests {
		t.Run(fmt.Sprintf("#%d", i), func(t *testing.T) {
			if got := part2Ofs(bytes.NewReader(tt.input)); got != tt.want {
				t.Errorf("ERROR: got: %v, want: %v", got, tt.want)
			}
		})
	}
}

func Benchmark_part2Ofs(b *testing.B) {
	for i := 0; i < b.N; i++ {
		if c := part2Ofs(bytes.NewReader(benchInput)); c != 3120 {
			b.Fatalf("ERROR: wrong offset: %d", c)
		}
	}
}
