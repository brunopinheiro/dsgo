package ds_test

import (
	"testing"

	"github.com/brunopinheiro/dsgo/ds"

	"github.com/stretchr/testify/require"
)

func TestList(t *testing.T) {
	fromArray := func(values []int) *ds.ListNode {
		var l *ds.ListNode
		for i, v := range values {
			if i == 0 {
				l = ds.NewListNode(v, nil)
				continue
			}

			l.Append(ds.NewListNode(v, nil))
		}
		return l
	}

	newNode := func(value int) *ds.ListNode {
		return ds.NewListNode(value, nil)
	}

	newList := func(value int, next *ds.ListNode) *ds.ListNode {
		return ds.NewListNode(value, next)
	}

	t.Run("appends to the end of the list", func(t *testing.T) {
		l := newNode(1)
		l.Append(newNode(2))
		l.Append(newNode(3))
		require.Equal(t, "1->2->3", l.Display())
	})

	t.Run("finds the kth element from the end", func(t *testing.T) {
		t.Run("returns not found when k is greater than the length of the list", func(t *testing.T) {
			_, found := ds.FindKthElementFromEnd(fromArray([]int{1, 2, 3, 4, 5}), 6)
			require.False(t, found)
		})

		t.Run("returns the first element when k is equal to the length of the list", func(t *testing.T) {
			value, found := ds.FindKthElementFromEnd(fromArray([]int{1, 2, 3, 4, 5}), 5)
			require.True(t, found)
			require.Equal(t, 1, value)
		})

		t.Run("returns the correct value when k is less than the length of the list", func(t *testing.T) {
			value, found := ds.FindKthElementFromEnd(fromArray([]int{1, 2, 3, 4, 5}), 3)
			require.True(t, found)
			require.Equal(t, 3, value)
		})

		t.Run("returns not found when k is 0", func(t *testing.T) {
			_, found := ds.FindKthElementFromEnd(fromArray([]int{1, 2, 3, 4, 5}), 0)
			require.False(t, found)
		})
	})

	t.Run("returns if list has a loop or not", func(t *testing.T) {
		t.Run("false when list is nil-terminated", func(t *testing.T) {
			require.False(t, ds.HasLoop(fromArray([]int{1, 2, 3, 4, 5})))
			require.False(t, ds.HasLoop(fromArray([]int{1})))
			require.False(t, ds.HasLoop(nil))
		})

		t.Run("true when list has a loop", func(t *testing.T) {
			loopPtr := newNode(3)
			endPtr := newList(5, loopPtr)
			loopPtr.Next = newList(4, endPtr)
			l := newList(1, newList(2, loopPtr))
			require.True(t, ds.HasLoop(l))
		})

		t.Run("single node loop", func(t *testing.T) {
			loopPtr := newNode(3)
			loopPtr.Next = loopPtr
			require.True(t, ds.HasLoop(loopPtr))
		})
	})

	t.Run("finds the beginning of the loop", func(t *testing.T) {
		t.Run("returns nil when list is nil-terminaned", func(t *testing.T) {
			require.Nil(t, ds.FindLoopStart(fromArray([]int{1, 2, 3, 4, 5})))
			require.Nil(t, ds.FindLoopStart(fromArray([]int{1})))
			require.Nil(t, ds.FindLoopStart(nil))
		})

		t.Run("the correct pointer for loops with length 1", func(t *testing.T) {
			loopPtr := newNode(3)
			loopPtr.Next = loopPtr
			require.Equal(t, loopPtr, ds.FindLoopStart(loopPtr))

			require.Equal(t, loopPtr, ds.FindLoopStart(
				newList(1, newList(2, loopPtr)),
			))
		})

		t.Run("the correct pointer for loops with length > 1", func(t *testing.T) {
			loopPtr := newNode(3)
			endPtr := newList(5, loopPtr)
			loopPtr.Next = newList(4, endPtr)
			l := newList(1, newList(2, loopPtr))

			require.Equal(t, loopPtr, ds.FindLoopStart(l))
		})
	})

	t.Run("finds the length of the loop", func(t *testing.T) {
		t.Run("returns 0 when the list is nil-terminated", func(t *testing.T) {
			require.Equal(t, 0, ds.LoopLength(fromArray([]int{1, 2, 3, 4, 5})))
			require.Equal(t, 0, ds.LoopLength(fromArray([]int{1})))
			require.Equal(t, 0, ds.LoopLength(nil))
		})

		t.Run("returns 1 when the list has a single node loop", func(t *testing.T) {
			loopPtr := newNode(3)
			loopPtr.Next = loopPtr
			require.Equal(t, 1, ds.LoopLength(loopPtr))

			require.Equal(t, 1, ds.LoopLength(
				newList(1, newList(2, loopPtr)),
			))
		})

		t.Run("the correct length for loops with length > 1", func(t *testing.T) {
			loopPtr := newNode(3)
			endPtr := newList(5, loopPtr)
			loopPtr.Next = newList(4, endPtr)
			l := newList(1, newList(2, loopPtr))

			require.Equal(t, 3, ds.LoopLength(l))
		})
	})

	t.Run("reverses the list", func(t *testing.T) {
		t.Run("empty list", func(t *testing.T) {
			require.Nil(t, ds.Reversed(nil))
		})

		t.Run("single node list", func(t *testing.T) {
			require.Equal(t, fromArray([]int{1}), ds.Reversed(newNode(1)))
		})

		t.Run("multiple node list", func(t *testing.T) {
			require.Equal(t, fromArray([]int{5, 4, 3, 2, 1}), ds.Reversed(
				fromArray([]int{1, 2, 3, 4, 5}),
			))
		})
	})
}
