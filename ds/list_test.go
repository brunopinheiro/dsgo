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
			_, found := fromArray([]int{1, 2, 3, 4, 5}).KthElementFromEnd(6)
			require.False(t, found)
		})

		t.Run("returns the first element when k is equal to the length of the list", func(t *testing.T) {
			value, found := fromArray([]int{1, 2, 3, 4, 5}).KthElementFromEnd(5)
			require.True(t, found)
			require.Equal(t, 1, value)
		})

		t.Run("returns the correct value when k is less than the length of the list", func(t *testing.T) {
			value, found := fromArray([]int{1, 2, 3, 4, 5}).KthElementFromEnd(3)
			require.True(t, found)
			require.Equal(t, 3, value)
		})

		t.Run("returns not found when k is 0", func(t *testing.T) {
			_, found := fromArray([]int{1, 2, 3, 4, 5}).KthElementFromEnd(0)
			require.False(t, found)
		})
	})

	t.Run("returns if list has a loop or not", func(t *testing.T) {
		t.Run("false when list is nil-terminated", func(t *testing.T) {
			require.False(t, fromArray([]int{1, 2, 3, 4, 5}).HasLoop())
			require.False(t, fromArray([]int{1}).HasLoop())
		})

		t.Run("true when list has a loop", func(t *testing.T) {
			loopPtr := newNode(3)
			endPtr := newList(5, loopPtr)
			loopPtr.Next = newList(4, endPtr)
			l := newList(1, newList(2, loopPtr))
			require.True(t, l.HasLoop())
		})

		t.Run("single node loop", func(t *testing.T) {
			loopPtr := newNode(3)
			loopPtr.Next = loopPtr
			require.True(t, loopPtr.HasLoop())
		})
	})

	t.Run("finds the beginning of the loop", func(t *testing.T) {
		t.Run("returns nil when list is nil-terminaned", func(t *testing.T) {
			require.Nil(t, fromArray([]int{1, 2, 3, 4, 5}).LoopStart())
			require.Nil(t, fromArray([]int{1}).LoopStart())
		})

		t.Run("the correct pointer for loops with length 1", func(t *testing.T) {
			loopPtr := newNode(3)
			loopPtr.Next = loopPtr
			require.Equal(t, loopPtr, loopPtr.LoopStart())

			require.Equal(t, loopPtr, newList(1, newList(2, loopPtr)).LoopStart())
		})

		t.Run("the correct pointer for loops with length > 1", func(t *testing.T) {
			loopPtr := newNode(3)
			endPtr := newList(5, loopPtr)
			loopPtr.Next = newList(4, endPtr)
			l := newList(1, newList(2, loopPtr))

			require.Equal(t, loopPtr, l.LoopStart())
		})
	})

	t.Run("finds the length of the loop", func(t *testing.T) {
		t.Run("returns 0 when the list is nil-terminated", func(t *testing.T) {
			require.Equal(t, 0, fromArray([]int{1, 2, 3, 4, 5}).LoopLength())
			require.Equal(t, 0, fromArray([]int{1}).LoopLength())
		})

		t.Run("returns 1 when the list has a single node loop", func(t *testing.T) {
			loopPtr := newNode(3)
			loopPtr.Next = loopPtr
			require.Equal(t, 1, loopPtr.LoopLength())

			require.Equal(t, 1, newList(1, newList(2, loopPtr)).LoopLength())
		})

		t.Run("the correct length for loops with length > 1", func(t *testing.T) {
			loopPtr := newNode(3)
			endPtr := newList(5, loopPtr)
			loopPtr.Next = newList(4, endPtr)
			l := newList(1, newList(2, loopPtr))

			require.Equal(t, 3, l.LoopLength())
		})
	})

	t.Run("reverses the list", func(t *testing.T) {
		t.Run("empty list", func(t *testing.T) {
			require.Nil(t, ds.ReverseList(nil))
		})

		t.Run("single node list", func(t *testing.T) {
			require.Equal(t, fromArray([]int{1}), ds.ReverseList(newNode(1)))
		})

		t.Run("multiple node list", func(t *testing.T) {
			require.Equal(t, fromArray([]int{5, 4, 3, 2, 1}), ds.ReverseList(
				fromArray([]int{1, 2, 3, 4, 5}),
			))
		})
	})

	t.Run("finds list merge point", func(t *testing.T) {
		t.Run("returns nil when there is no merge point", func(t *testing.T) {
			require.Nil(t, ds.FindListsMergePoint(
				fromArray([]int{1, 2, 3}),
				fromArray([]int{4, 5, 6}),
			))

			require.Nil(t, ds.FindListsMergePoint(
				nil,
				fromArray([]int{4, 5, 6}),
			))

			require.Nil(t, ds.FindListsMergePoint(
				fromArray([]int{1, 2, 3}),
				nil,
			))
		})

		t.Run("when there's a merge point", func(t *testing.T) {
			sharedList := fromArray([]int{7, 8, 9})

			t.Run("and one of the lists starts at the merge point", func(t *testing.T) {
				require.Equal(
					t,
					sharedList,
					ds.FindListsMergePoint(
						sharedList,
						newList(5, newList(6, sharedList)),
					),
				)

				require.Equal(
					t,
					sharedList,
					ds.FindListsMergePoint(
						newList(5, newList(6, sharedList)),
						sharedList,
					),
				)
			})

			t.Run("and lists have the same length", func(t *testing.T) {
				require.Equal(
					t,
					sharedList,
					ds.FindListsMergePoint(
						newList(1, newList(2, sharedList)),
						newList(5, newList(6, sharedList)),
					),
				)
			})

			t.Run("and lists have different lengths", func(t *testing.T) {
				require.Equal(
					t,
					sharedList,
					ds.FindListsMergePoint(
						newList(1, newList(2, newList(3, newList(4, sharedList)))),
						newList(5, newList(6, sharedList)),
					),
				)

				require.Equal(
					t,
					sharedList,
					ds.FindListsMergePoint(
						newList(4, newList(5, newList(6, sharedList))),
						newList(1, sharedList),
					),
				)
			})
		})
	})

	t.Run("middle of the list", func(t *testing.T) {
		t.Run("with an odd number of elements", func(t *testing.T) {
			require.Equal(t, 3, fromArray([]int{1, 2, 3, 4, 5}).Middle())
		})

		t.Run("with an even number of elements", func(t *testing.T) {
			require.Equal(t, 4, fromArray([]int{1, 2, 3, 4, 5, 6}).Middle())
		})
	})

	t.Run("reverse display", func(t *testing.T) {
		l := fromArray([]int{1, 2, 3, 4, 5})
		require.Equal(t, "5->4->3->2->1", l.ReverseDisplay())
		require.Equal(t, "1->2->3->4->5", l.Display()) // it does not alter the list

		require.Equal(t, "1", fromArray([]int{1}).ReverseDisplay())
	})
}
