package main

import (
	"bytes"
	"fmt"
	"testing"
)

func Test_forEachElf(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		wantCnt  int
		wantMax  int
		wantLast int
	}{
		{input: "1000\n2000\n3000\n\n4000\n\n5000\n6000\n\n7000\n8000\n9000\n\n10000", wantCnt: 5, wantMax: 24000, wantLast: 10000},
	}
	for i, tt := range tests {
		t.Run(fmt.Sprintf("#%d", i), func(t *testing.T) {
			var err error
			var cnt, max, last int
			err = forEachElf(bytes.NewBufferString(tt.input), func(cals []int) {
				cnt++
				var sum int
				for _, cal := range cals {
					sum += cal
				}
				if sum > max {
					max = sum
				}
				last = sum
			})
			if err != nil {
				t.Fatalf("ERROR: forEachElf failed with error: %v", err)
			}
			if cnt != tt.wantCnt {
				t.Errorf("ERROR: got cnt: %d, want: %d", cnt, tt.wantCnt)
			}

			if max != tt.wantMax {
				t.Errorf("ERROR: got max: %d, want: %d", max, tt.wantMax)
			}

			if last != tt.wantLast {
				t.Errorf("ERROR: got last: %d, want: %d", last, tt.wantLast)
			}
		})
	}
}
