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

/*
Why does it work? (for lists with cycle)
- lets assume the distance from head to cycle beginning is n
- lets assume the length of the cycle is L

- when the slow pointer moves n, the faster pointer will have moved 2n
- if we consider the cycle beginning to be at position 0:
  - the slow pointer will be at position 0
  - faster pointer will be at position n
  - for the faster pointer to reach the cycle beginning again, it must move (L - n)
  - we want to find x, where 0 + x <=> (L - n) + 2x
  - so x = L - n + 2x =>
  - x + n = L + 2x =>
  - n = L + x
  - since L is a complete cycle, safe to assume n = x
  - this means slow and fast pointers will meet at position n

- moving slow back to the head (position = -n)
  - fast pointer will be at position n => (L - n) from the cycle beginning
  - this time, moving both pointers 1 by 1, after n steps:
  - slow pointer will be at position 0
  - fast pointer will be at position (L - n) + n => L <=> 0
*/
func FindBeginningOfCycle(l *ListNode) *ListNode {
	if l == nil {
		return nil
	}

	slowPtr := l
	fastPtr := l
	hasCycle := false

	for fastPtr.Next != nil && fastPtr.Next.Next != nil {
		slowPtr = slowPtr.Next
		fastPtr = fastPtr.Next.Next
		if slowPtr == fastPtr {
			hasCycle = true
			break
		}
	}

	if !hasCycle {
		return nil
	}

	slowPtr = l
	for slowPtr != fastPtr {
		slowPtr = slowPtr.Next
		fastPtr = fastPtr.Next
	}

	return slowPtr
}
