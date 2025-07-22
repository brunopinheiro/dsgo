package algorithms_test

import (
	"testing"

	"github.com/brunopinheiro/dsgo/algorithms"
	"github.com/stretchr/testify/require"
)

func TestMaxNumberInSlidingWindows(t *testing.T) {
	t.Run("when window size equal to array length", func(t *testing.T) {
		require.Equal(
			t,
			[]int{5},
			algorithms.MaxNumberInSlidingWindows(
				[]int{1, 5, 3},
				3,
			),
		)
	})

	t.Run("when window size is smaller than array length", func(t *testing.T) {
		require.Equal(
			t,
			[]int{5, 5},
			algorithms.MaxNumberInSlidingWindows(
				[]int{1, 5, 3, 4},
				3,
			),
		)
	})

	t.Run("when array length is at least twice the size of the window", func(t *testing.T) {
		require.Equal(
			t,
			[]int{5, 5, 7, 7, 7, 9},
			algorithms.MaxNumberInSlidingWindows(
				[]int{1, 5, 3, 4, 7, -1, 0, 9},
				3,
			),
		)
	})
}
