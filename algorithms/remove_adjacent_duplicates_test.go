package algorithms_test

import (
	"fmt"
	"testing"

	"github.com/brunopinheiro/dsgo/algorithms"
	"github.com/stretchr/testify/require"
)

func TestRemoveAdjacentDuplicates(t *testing.T) {
	testCases := map[string]string{
		"abc":         "abc",
		"aabc":        "bc",
		"aaabc":       "abc",
		"abbc":        "ac",
		"abcc":        "ab",
		"abbac":       "c",
		"abbaac":      "ac",
		"career":      "ca",
		"mississippi": "m",
		"aaaaa":       "a",
		"acacc":       "aca",
		"acaacc":      "ac",
	}

	for input, expected := range testCases {
		t.Run(fmt.Sprintf("%s => %s", input, expected), func(t *testing.T) {
			require.Equal(t, expected, algorithms.RemoveAdjacentDuplicates(input))
		})
	}
}
