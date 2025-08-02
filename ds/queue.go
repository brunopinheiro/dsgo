package ds

import (
	"fmt"
	"strings"
)

// for queue implementation, decided to go with a dynamic array
// maybe in future I can try implementing a pointer version for queue
// and dynamic array implementations for the other structures

type Queue struct {
	values []int
}

func NewQueue() *Queue {
	return &Queue{values: []int{}}
}

func (q *Queue) Size() int {
	return len(q.values)
}

func (q *Queue) Enqueue(value int) {
	q.values = append(q.values, value)
}

func (q *Queue) Dequeue() int {
	if len(q.values) == 0 {
		panic("queue is empty")
	}
	v := q.values[0]
	q.values = q.values[1:]
	return v
}

func (q *Queue) Display() string {
	rawValues := []string{}
	for _, v := range q.values {
		rawValues = append(rawValues, fmt.Sprintf("%d", v))
	}
	return strings.Join(rawValues, "->")
}

func (q *Queue) Reverse() {
	if len(q.values) == 0 {
		return
	}

	leftPtr := 0
	rightPtr := len(q.values) - 1
	for rightPtr > leftPtr {
		q.values[leftPtr] = q.values[leftPtr] ^ q.values[rightPtr]
		q.values[rightPtr] = q.values[leftPtr] ^ q.values[rightPtr]
		q.values[leftPtr] = q.values[leftPtr] ^ q.values[rightPtr]
		leftPtr += 1
		rightPtr -= 1
	}
}

func QueueMergeHalfs(q *Queue) *List {
	var l *List
	rightHalfQueue := NewQueue()
	for range (q.Size() + 1) / 2 {
		rightHalfQueue.Enqueue(q.Dequeue())
	}
	for range rightHalfQueue.Size() {
		l = ListAppend(l, NewList(rightHalfQueue.Dequeue(), nil))
		if q.Size() > 0 {
			l = ListAppend(l, NewList(q.Dequeue(), nil))
		}
	}
	return l
}
