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

	t.Run("returns if list has a cycle or not", func(t *testing.T) {
		t.Run("false when list is nil-terminated", func(t *testing.T) {
			require.False(t, ds.HasCycle(fromArray([]int{1, 2, 3, 4, 5})))
			require.False(t, ds.HasCycle(fromArray([]int{1})))
			require.False(t, ds.HasCycle(nil))
		})

		t.Run("true when list has a cycle", func(t *testing.T) {
			cyclePtr := newNode(3)
			endPtr := newList(5, cyclePtr)
			cyclePtr.Next = newList(4, endPtr)
			l := newList(1, newList(2, cyclePtr))
			require.True(t, ds.HasCycle(l))
		})

		t.Run("single node cycle", func(t *testing.T) {
			cyclePtr := newNode(3)
			cyclePtr.Next = cyclePtr
			require.True(t, ds.HasCycle(cyclePtr))
		})
	})

	t.Run("finds the beginning of the cycle", func(t *testing.T) {
		t.Run("returns nil when list is nil-terminaned", func(t *testing.T) {
			require.Nil(t, ds.FindBeginningOfCycle(fromArray([]int{1, 2, 3, 4, 5})))
			require.Nil(t, ds.FindBeginningOfCycle(fromArray([]int{1})))
			require.Nil(t, ds.FindBeginningOfCycle(nil))
		})

		t.Run("true when list has a cycle", func(t *testing.T) {
			cyclePtr := newNode(3)
			endPtr := newList(5, cyclePtr)
			cyclePtr.Next = newList(4, endPtr)
			l := newList(1, newList(2, cyclePtr))

			require.Equal(t, cyclePtr, ds.FindBeginningOfCycle(l))
		})

		t.Run("single node cycle", func(t *testing.T) {
			cyclePtr := newNode(3)
			cyclePtr.Next = cyclePtr
			require.Equal(t, cyclePtr, ds.FindBeginningOfCycle(cyclePtr))
		})
	})
}
