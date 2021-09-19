package main

import "testing"

type Input struct {
	s       string
	numRows int
}

func TestConvert(t *testing.T) {
	inputs := []Input{
		{s: "PAYPALISHIRING", numRows: 3},
		{s: "PAYPALISHIRING", numRows: 4},
	}
	outputs := []string{"PAHNAPLSIIGYIR", "PINALSIGYAHRPI"}

	for i, input := range inputs {
		result := Convert(input.s, input.numRows)

		if result != outputs[i] {
			t.Errorf("Solution %d is not correct\n", i)
		}
	}
}
