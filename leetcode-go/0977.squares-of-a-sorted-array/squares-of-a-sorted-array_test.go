package p0977

import (
	"reflect"
	"testing"
)

var tests = map[string]struct {
	input []int
	want []int
}{
	"均为正数或0": {
		[]int{0, 1, 2, 3, 10},
		[]int{0, 1, 4, 9, 100},
	},
	"均为负数或0": {
		[]int{-7, -3, -2, -1, 0},
		[]int{0, 1, 4, 9, 49},
	},
	"正数、负数混合": {
		[]int{-7, -3, 2, 3, 11},
		[]int{4, 9, 9, 49, 121},
	},
}

func TestSortedSquares(t *testing.T) {
	for name, tc := range tests {
		t.Run(name, func(t *testing.T) {
			got := SortedSquares(tc.input)
			if !reflect.DeepEqual(tc.want, got) {
				t.Fatalf("%s expected: %#v, got: %#v", name, tc.want, got)
			}
		})
	}
}

func Benchmark_sortedSquares(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, tc := range tests {
			SortedSquares(tc.input)
		}
	}
}
