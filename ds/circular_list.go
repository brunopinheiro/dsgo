package ds

import (
	"fmt"
	"strings"
)

type CircularListNode struct {
	Value int
	Next  *CircularListNode
}

func NewCircularList(value int) *CircularListNode {
	l := &CircularListNode{Value: value, Next: nil}
	l.Next = l
	return l
}

func (l *CircularListNode) Display() string {
	values := []string{
		fmt.Sprintf("%d", l.Value),
	}
	cursor := l.Next
	for cursor != l {
		values = append(values, fmt.Sprintf("%d", cursor.Value))
		cursor = cursor.Next
	}
	return strings.Join(values, "->")
}

func (l *CircularListNode) Append(value int) {
	cursor := l
	for cursor.Next != l {
		cursor = cursor.Next
	}

	cursor.Next = &CircularListNode{Value: value, Next: l}
}

func (l *CircularListNode) JosephusCircle(k int) int {
	if k < 1 {
		panic("k must be greater than 0")
	}

	// to be able to remove the first player and keep
	// the counting consistent, we need to start prev at the last element
	// not sure if there's something we can do here to avoid this extra step
	// :thinking_face: maybe create a wrapper structure that can keep a reference to the last node
	prev := l
	cursor := l.Next
	for cursor != l {
		prev = prev.Next
		cursor = cursor.Next
	}

	for cursor != prev {
		for range k - 1 {
			prev = prev.Next
			cursor = cursor.Next
		}
		prev.Next = cursor.Next
		cursor = cursor.Next
	}

	return cursor.Value
}

func SplitCircularListInHalf(l *CircularListNode) (*CircularListNode, *CircularListNode) {
	if l == nil {
		return nil, nil
	}

	if l.Next == l {
		return l, nil
	}

	slowPtr := l
	fastPtr := l
	for fastPtr.Next != l && fastPtr.Next.Next != l {
		slowPtr = slowPtr.Next
		fastPtr = fastPtr.Next.Next
	}

	right := slowPtr.Next
	slowPtr.Next = l

	if fastPtr.Next == l {
		fastPtr.Next = right
	} else {
		fastPtr.Next.Next = right
	}

	return l, right
}
