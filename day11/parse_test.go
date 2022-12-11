package main

import (
	_ "embed"
	"fmt"
	"reflect"
	"testing"

	"aoc2022/pkg/ilr"
)

//go:embed test.txt
var testInput string

//go:embed input.txt
var benchInput string

func Test_line2items(t *testing.T) {
	tests := []struct {
		line      string
		wantItems []int
	}{
		{line: "  Starting items: 79, 98", wantItems: []int{79, 98}},
		{line: "  Starting items: 54, 65, 75, 74", wantItems: []int{54, 65, 75, 74}},
	}
	for i, tt := range tests {
		t.Run(fmt.Sprintf("#%d", i), func(t *testing.T) {
			if gotItems := line2items(tt.line); !reflect.DeepEqual(gotItems, tt.wantItems) {
				t.Errorf("ERROR: got:\n%v\nwant:\n%v", gotItems, tt.wantItems)
			}
		})
	}
}

func Test_line2op(t *testing.T) {
	tests := []struct {
		line string
		want int
	}{
		{line: "  Operation: new = old * 19", want: 2 * 19},
		{line: "  Operation: new = old + 6", want: 2 + 6},
		{line: "  Operation: new = old * old", want: 2 * 2},
		{line: "  Operation: new = old + old", want: 2 + 2},
	}
	for i, tt := range tests {
		t.Run(fmt.Sprintf("#%d", i), func(t *testing.T) {
			got := line2op(tt.line)(2)
			if got != tt.want {
				t.Errorf("ERROR: got: %d, want: %d", got, tt.want)
			}
		})
	}
}

func Test_line2int(t *testing.T) {
	tests := []struct {
		line string
		want int
	}{
		{line: "  Test: divisible by 5", want: 5},
		{line: "  Test: divisible by 23", want: 23},
		{line: "    If true: throw to monkey 2", want: 2},
		{line: "    If false: throw to monkey 3", want: 3},
	}
	for i, tt := range tests {
		t.Run(fmt.Sprintf("#%d", i), func(t *testing.T) {
			if got := line2int(tt.line); got != tt.want {
				t.Errorf("ERROR: got: %d, want: %d", got, tt.want)
			}
		})
	}
}

func Test_part1Score(t *testing.T) {
	tests := []struct {
		input string
		want  int
	}{
		{input: testInput, want: 10605},
		{input: benchInput, want: 58786},
	}
	for i, tt := range tests {
		t.Run(fmt.Sprintf("#%d", i), func(t *testing.T) {
			if got := part1Score(ilr.New(tt.input)); got != tt.want {
				t.Errorf("ERROR: got: %d, want: %d", got, tt.want)
			}
		})
	}
}

func Benchmark_part1Score(b *testing.B) {
	for i := 0; i < b.N; i++ {
		if c := part1Score(ilr.New(benchInput)); c != 58786 {
			b.Fatalf("ERROR: wrong number: %d", c)
		}
	}
}

func Test_part2Rounds(t *testing.T) {
	monkeys := parse(ilr.New(testInput))
	monkeys.calcLcm()
	tests := []struct {
		rounds int
		want   []int
	}{
		{rounds: 1, want: []int{2, 4, 3, 6}},                   //     1 rounds
		{rounds: 19, want: []int{99, 97, 8, 103}},              //    20 rounds
		{rounds: 980, want: []int{5204, 4792, 199, 5192}},      //  1000 rounds
		{rounds: 1000, want: []int{10419, 9577, 392, 10391}},   //  2000 rounds
		{rounds: 1000, want: []int{15638, 14358, 587, 15593}},  //  3000 rounds
		{rounds: 1000, want: []int{20858, 19138, 780, 20797}},  //  4000 rounds
		{rounds: 1000, want: []int{26075, 23921, 974, 26000}},  //  5000 rounds
		{rounds: 1000, want: []int{31294, 28702, 1165, 31204}}, //  6000 rounds
		{rounds: 1000, want: []int{36508, 33488, 1360, 36400}}, //  7000 rounds
		{rounds: 1000, want: []int{41728, 38268, 1553, 41606}}, //  8000 rounds
		{rounds: 1000, want: []int{46945, 43051, 1746, 46807}}, //  9000 rounds
		{rounds: 1000, want: []int{52166, 47830, 1938, 52013}}, // 10000 rounds
	}
	for i, tt := range tests {
		t.Run(fmt.Sprintf("#%d", i), func(t *testing.T) {
			monkeys.play(2, tt.rounds)
			for i, want := range tt.want {
				if monkeys[i].ic != want {
					t.Errorf("ERROR: got (monkey #%d): %d, want: %d", i, monkeys[i].ic, want)
				}
			}
		})
	}
}

func Test_part2Score(t *testing.T) {
	tests := []struct {
		input string
		want  int
	}{
		{input: testInput, want: 2713310158},
		{input: benchInput, want: 14952185856},
	}
	for i, tt := range tests {
		t.Run(fmt.Sprintf("#%d", i), func(t *testing.T) {
			if got := part2Score(ilr.New(tt.input)); got != tt.want {
				t.Errorf("ERROR: got: %d, want: %d", got, tt.want)
			}
		})
	}
}

func Benchmark_part2Score(b *testing.B) {
	for i := 0; i < b.N; i++ {
		if c := part2Score(ilr.New(benchInput)); c != 14952185856 {
			b.Fatalf("ERROR: wrong number: %d", c)
		}
	}
}
