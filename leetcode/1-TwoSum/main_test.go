package main

import (
	"reflect"
	"testing"
)

func TestTwoSum(t *testing.T) {
	nums := []int{2, 7, 11, 15}
	target := 9
	res := TwoSum(nums, target)
	expected := []int{0, 1}

	if !reflect.DeepEqual(res, expected) {
		t.Error("Solution was incorrect")
	}
}
