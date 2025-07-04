package ds

import (
	"fmt"
	"strings"
)

type ListNode struct {
	Value int
	Next  *ListNode
}

func NewListNode(value int, next *ListNode) *ListNode {
	return &ListNode{Value: value, Next: next}
}

func (n *ListNode) Append(newNode *ListNode) {
	cursor := n
	for cursor.Next != nil {
		cursor = cursor.Next
	}
	cursor.Next = newNode
}

func (n *ListNode) Display() string {
	values := []string{}
	cursor := n
	for cursor != nil {
		values = append(values, fmt.Sprintf("%d", cursor.Value))
		cursor = cursor.Next
	}
	return strings.Join(values, "->")
}

func FindKthElementFromEnd(l *ListNode, k uint) (int, bool) {
	firstPtr := l
	kthPtr := l

	if k == 0 {
		return 0, false
	}

	for range k {
		if firstPtr == nil {
			return 0, false
		}
		firstPtr = firstPtr.Next
	}

	for firstPtr != nil {
		firstPtr = firstPtr.Next
		kthPtr = kthPtr.Next
	}

	return kthPtr.Value, true
}

func HasCycle(l *ListNode) bool {
	if l == nil {
		return false
	}

	slowPtr := l
	fastPtr := l

	for fastPtr.Next != nil && fastPtr.Next.Next != nil {
		slowPtr = slowPtr.Next
		fastPtr = fastPtr.Next.Next
		if slowPtr == fastPtr {
			return true
		}
	}
	return false
}
