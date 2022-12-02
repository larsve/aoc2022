package main

import (
	"bytes"
	"testing"
)

func Test_scorePart1(t *testing.T) {
	tests := []struct {
		name      string
		input     string
		wantScore int
	}{
		{input: "A Y\nB X\nC Z", wantScore: 15},
		{input: "A Z\nB X\nC Y", wantScore: 6},
		{input: "A X\nB Y\nC Z", wantScore: 15},
		{input: "A Y\nB Z\nC X", wantScore: 24},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := scorePart1(bytes.NewBufferString(tt.input)); got != tt.wantScore {
				t.Errorf("ERROR: got score: %v, want %v", got, tt.wantScore)
			}
		})
	}
}

func Test_scorePart2(t *testing.T) {
	tests := []struct {
		name      string
		input     string
		wantScore int
	}{
		{input: "A Y\nB X\nC Z", wantScore: 12},
		{input: "A X\nB X\nC X", wantScore: 6},
		{input: "A Y\nB Y\nC Y", wantScore: 15},
		{input: "A Z\nB Z\nC Z", wantScore: 24},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := scorePart2(bytes.NewBufferString(tt.input)); got != tt.wantScore {
				t.Errorf("ERROR: got score: %v, want %v", got, tt.wantScore)
			}
		})
	}
}
