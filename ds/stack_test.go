package ds_test

import (
	"testing"

	"github.com/brunopinheiro/dsgo/ds"
	"github.com/stretchr/testify/require"
)

func TestStack(t *testing.T) {
	fromArray := func(values []int) *ds.Stack {
		var s *ds.Stack
		for _, v := range values {
			s = ds.StackPush(s, v)
		}
		return s
	}

	t.Run("pushes values to the top of the stack", func(t *testing.T) {
		s := ds.NewStack(1)
		s = ds.StackPush(s, 2)
		s = ds.StackPush(s, 3)
		require.Equal(t, "3->2->1", s.Display())
	})

	t.Run("pops remove the top of the stack", func(t *testing.T) {
		s := fromArray([]int{1, 2, 3, 4})
		s, value := ds.StackPop(s)
		require.Equal(t, 4, value)
		require.Equal(t, "3->2->1", s.Display())
	})

	t.Run("IsArrayPalindrome", func(t *testing.T) {
		// function will receive an array of integers
		// 0 will indicate the middle of the array
		// we can assume all other values are > 0

		t.Run("given array is empty", func(t *testing.T) {
			require.True(t, ds.IsArrayPalindrome([]int{}))
		})

		t.Run("given array as a single element", func(t *testing.T) {
			require.True(t, ds.IsArrayPalindrome([]int{1}))
		})

		t.Run("given array is a palindrome", func(t *testing.T) {
			require.True(t, ds.IsArrayPalindrome([]int{1, 2, 3, 0, 3, 2, 1}))
		})

		t.Run("given array is not a palindrome", func(t *testing.T) {
			require.False(t, ds.IsArrayPalindrome([]int{1, 2, 3, 0, 3, 4, 1}))
		})

		t.Run("given array has an even number of elements", func(t *testing.T) {
			require.False(t, ds.IsArrayPalindrome([]int{1, 2, 3, 0, 3, 2}))
			require.False(t, ds.IsArrayPalindrome([]int{1, 2, 0, 2, 1, 1}))
		})
	})
}
