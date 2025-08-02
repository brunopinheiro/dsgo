package ds_test

import (
	"testing"

	"github.com/brunopinheiro/dsgo/ds"
	"github.com/stretchr/testify/require"
)

func TestQueue(t *testing.T) {
	fromArray := func(values []int) *ds.Queue {
		q := ds.NewQueue()
		for _, v := range values {
			q.Enqueue(v)
		}
		return q
	}

	t.Run("enqueue elements", func(t *testing.T) {
		q := ds.NewQueue()
		q.Enqueue(1)
		q.Enqueue(2)
		q.Enqueue(3)
		require.Equal(t, "1->2->3", q.Display())
	})

	t.Run("dequeue elements", func(t *testing.T) {
		q := fromArray([]int{1, 2, 3, 4, 5})
		actual := q.Dequeue()
		require.Equal(t, 1, actual)

		actual = q.Dequeue()
		require.Equal(t, 2, actual)

		require.Equal(t, "3->4->5", q.Display())
	})

	t.Run("reverse", func(t *testing.T) {
		q := fromArray([]int{1, 2, 3, 4, 5})
		q.Reverse()
		require.Equal(t, "5->4->3->2->1", q.Display())
	})

	t.Run("merge halfs", func(t *testing.T) {
		t.Run("empty queue", func(t *testing.T) {
			require.Nil(t, ds.QueueMergeHalfs(ds.NewQueue()))
		})

		t.Run("single element", func(t *testing.T) {
			q := ds.NewQueue()
			q.Enqueue(1)
			require.Equal(t, "1", ds.QueueMergeHalfs(q).Display())
		})

		t.Run("odd number of elements", func(t *testing.T) {
			q := fromArray([]int{1, 2, 3, 4, 5})
			require.Equal(t, "1->4->2->5->3", ds.QueueMergeHalfs(q).Display())
		})

		t.Run("even number of elements", func(t *testing.T) {
			q := fromArray([]int{1, 2, 3, 4, 5, 6})
			require.Equal(t, "1->4->2->5->3->6", ds.QueueMergeHalfs(q).Display())
		})
	})

	t.Run("check sequence in pairs", func(t *testing.T) {
		t.Run("empty queue", func(t *testing.T) {
			require.True(t, ds.QueueCheckSequenceInPairs(ds.NewQueue()))
		})

		t.Run("single element", func(t *testing.T) {
			q := ds.NewQueue()
			q.Enqueue(1)
			require.True(t, ds.QueueCheckSequenceInPairs(q))
		})

		t.Run("odd number of elements", func(t *testing.T) {
			q := fromArray([]int{99, -1, -2, 6, 7})
			require.True(t, ds.QueueCheckSequenceInPairs(q))
		})

		t.Run("even number of elements", func(t *testing.T) {
			q := fromArray([]int{99, 98, -1, -2, 6, 7})
			require.True(t, ds.QueueCheckSequenceInPairs(q))
		})

		t.Run("when a pair is not in sequence", func(t *testing.T) {
			q := fromArray([]int{99, 98, -1, -2, 15, 17, 6, 7})
			require.False(t, ds.QueueCheckSequenceInPairs(q))
		})
	})
}
