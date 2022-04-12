package test

import (
	"testing"
)

type ProducePiOfTest struct {
	Input    []int
	Expected []int
}

var ProducePiOfTests = []ProducePiOfTest{
	{Input: []int{2, 1, 3, 4, 5}, Expected: []int{60, 120, 40, 30, 24}},
	{Input: []int{2, 1, 0, 4, 5}, Expected: []int{0, 0, 40, 0, 0}},
	{Input: []int{2, 1, 0, 4, 0}, Expected: []int{0, 0, 0, 0, 0}},
}

func TestProducePiOf(t *testing.T) {
	for _, test := range ProducePiOfTests {
		if output := ProducePiOf(test.Input); !Equal(output, test.Expected) {
			t.Errorf("Output %v not equal to expected %v", output, test.Expected)
		}
	}
}

func Equal(a, b []int) bool {
	if len(a) != len(b) {
		return false
	}
	for i, v := range a {
		if v != b[i] {
			return false
		}
	}
	return true
}
