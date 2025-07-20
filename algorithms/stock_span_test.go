package algorithms_test

import (
	"testing"

	"github.com/brunopinheiro/dsgo/algorithms"
	"github.com/stretchr/testify/require"
)

func TestStockSpan(t *testing.T) {
	t.Run("for an empty array", func(t *testing.T) {
		require.Empty(
			t,
			algorithms.StockSpan([]int{}),
		)
	})

	t.Run("for an array with single element", func(t *testing.T) {
		require.Equal(
			t,
			[]int{1},
			algorithms.StockSpan([]int{99}),
		)
	})

	t.Run("for an array with decreasing values", func(t *testing.T) {
		require.Equal(
			t,
			[]int{1, 1, 1, 1, 1},
			algorithms.StockSpan([]int{5, 4, 3, 2, 1}),
		)
	})

	t.Run("for an array with increasing values", func(t *testing.T) {
		require.Equal(
			t,
			[]int{1, 2, 3, 4, 5},
			algorithms.StockSpan([]int{1, 2, 3, 4, 5}),
		)
	})

	t.Run("for an array with mixed values", func(t *testing.T) {
		require.Equal(
			t,
			[]int{1, 1, 1, 2, 1, 4, 6},
			algorithms.StockSpan([]int{100, 80, 60, 70, 60, 75, 85}),
		)
	})
}
