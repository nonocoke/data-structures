package p0078

import (
	"reflect"
	"testing"
)

var tests = map[string]struct {
	input []int
	want [][]int
}{

	"nums为空": {
	 	[]int{},
	 	[][]int{nil},
	},
	"case 1": {
		[]int{1, 2, 3},
		[][]int{
			nil,
			{1},
			{2},
			{1, 2},
			{3},
			{1, 3},
			{2, 3},
			{1, 2, 3},
		},
	},
	"case 2": {
		[]int{1, 2, 3, 4},
		[][]int{
			nil,
			{1},
			{2},
			{1, 2},
			{3},
			{1, 3},
			{2, 3},
			{1, 2, 3},
			{4},
			{1, 4},
			{2, 4},
			{1, 2, 4},
			{3, 4},
			{1, 3, 4},
			{2, 3, 4},
			{1, 2, 3, 4},
		},
	},
	"case 3": {  // Error
		[]int{9, 0, 3, 5, 7},
		[][]int{nil},
	},
}

func TestSubsets(t *testing.T) {
	for name, tc := range tests {
		t.Run(name, func(t *testing.T) {
			got := Subsets(tc.input)
			if !reflect.DeepEqual(tc.want, got) {
				t.Fatalf("%s expected: %#v, got: %#v", name, tc.want, got)
			}
		})
	}
}
