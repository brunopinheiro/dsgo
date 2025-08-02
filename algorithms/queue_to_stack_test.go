package algorithms_test

import (
	"testing"

	"github.com/brunopinheiro/dsgo/algorithms"
	"github.com/brunopinheiro/dsgo/ds"
	"github.com/stretchr/testify/require"
)

func TestCopyQueueOntoStackTest(t *testing.T) {
	t.Run("when queue is empty", func(t *testing.T) {
		q := ds.NewQueue()
		s := algorithms.QueueToStack(q)
		require.Equal(t, "", s.Display())
	})

	t.Run("when queue has a single value", func(t *testing.T) {
		q := ds.NewQueue()
		q.Enqueue(1)
		s := algorithms.QueueToStack(q)
		require.Equal(t, "1", s.Display())
	})

	t.Run("when queue has multiple values", func(t *testing.T) {
		q := ds.NewQueue()
		q.Enqueue(1)
		q.Enqueue(2)
		q.Enqueue(3)
		s := algorithms.QueueToStack(q)
		require.Equal(t, "1->2->3", s.Display())
	})
}
