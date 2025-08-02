package algorithms

import (
	"github.com/brunopinheiro/dsgo/ds"
)

// challenge is to only use enqueue, dequeue, push and pop
func QueueToStack(q *ds.Queue) *ds.Stack {
	if q.Size() == 0 {
		return nil
	}

	var s *ds.Stack
	for q.Size() != 0 {
		value := q.Dequeue()
		s = ds.StackPush(s, value)
	}

	for s != nil {
		newStack, value := ds.StackPop(s)
		s = newStack
		q.Enqueue(value)
	}

	for q.Size() != 0 {
		value := q.Dequeue()
		s = ds.StackPush(s, value)
	}

	return s
}
