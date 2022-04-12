package test

import (
	"testing"
)

type PaintMatrixWith1Test struct {
	Input    [][]int
	Expected [][]int
}

var PaintMatrixWith1Tests = []PaintMatrixWith1Test{
	{
		Input: [][]int{
			{0, 1, 2, 3},
			{3, 1, 2, 4},
			{1, 0, 2, 3},
			{5, 9, 2, 5},
		},
		Expected: [][]int{
			{1, 1, 1, 1},
			{1, 1, 1, 1},
			{1, 1, 1, 1},
			{1, 1, 2, 5},
		},
	},
}

func TestPaintMatrixWith1(t *testing.T) {
	for _, test := range PaintMatrixWith1Tests {
		if output := PaintMatrixWith1(test.Input); !MatrixEqual(output, test.Expected) {
			t.Errorf("Output %v not equal to expected %v", output, test.Expected)
		}
	}
}

func MatrixEqual(m1, m2 [][]int) bool {
	if len(m1) != len(m2) || len(m1[0]) != len(m2[0]) {
		return false
	} else {
		for i := 0; i < len(m1); i++ {
			for j := 0; j < len(m1[0]); j++ {
				if m1[i][j] != m2[i][j] {
					return false
				}
			}
		}
	}
	return true
}
