package test

import (
	"testing"
)

type SpiralPatternTest struct {
	Input    int
	Expected string
}

var SpiralPatternTests = []SpiralPatternTest{
	{2, "1 2 4 3"},
	{3, "1 2 3 6 9 8 7 4 5"},
	{4, "1 2 3 4 8 12 16 15 14 13 9 5 6 7 11 10"},
}

func TestSpiralPattern(t *testing.T) {
	for _, test := range SpiralPatternTests {
		if output := SpiralPattern(test.Input); output != test.Expected {
			t.Errorf("Output %s not equal to expected %s", output, test.Expected)
		}
	}
}
