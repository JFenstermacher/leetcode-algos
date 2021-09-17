package main

import (
	"fmt"
	"sort"
	"testing"
)

func compare(a [][]string, b [][]string) bool {
	if len(a) != len(b) {
		return false
	}

	mapping := map[string]interface{}{}

	for _, arr := range a {
		sort.Strings(arr)

		key := fmt.Sprint(arr)

		mapping[key] = nil
	}

	for _, arr := range b {
		sort.Strings(arr)

		key := fmt.Sprint(arr)

		_, found := mapping[key]

		if !found {
			return false
		}
	}

	return true
}

func TestGroupAnagrams(t *testing.T) {
	input := []string{"eat", "tea", "tan", "ate", "nat", "bat"}
	expected := [][]string{
		{"bat"},
		{"nat", "tan"},
		{"ate", "eat", "tea"},
	}

	output := groupAnagrams(input)

	if !compare(output, expected) {
		t.Error("Failed test case")
	}
}
