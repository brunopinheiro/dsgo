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
}
