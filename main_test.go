package main

import (
	"fmt"
	"strings"
	"testing"
)

func TestConstruct(t *testing.T) {
	var tests = []struct {
		caps     map[int]int
		tables   map[int]struct{}
		n        int
		expected string
	}{
		{map[int]int{}, map[int]struct{}{}, 6, strings.Repeat(" ", 6)},
		{map[int]int{0: 1, 2: 3}, map[int]struct{}{1: struct{}{}}, 6, ".T:.  "},
	}

	for _, test := range tests {
		if output := construct(test.caps, test.tables, test.n); output != test.expected {
			t.Errorf("Test Failed: {%s} inputted, {%s} expected, recieved: {%s}\n",
				fmt.Sprint(test.caps, test.tables, test.n), test.expected, output)
		}
	}

}
