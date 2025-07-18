package ds

import (
	"fmt"
	"strings"
)

type List struct {
	Value int
	Next  *List
}

func NewList(value int, next *List) *List {
	return &List{Value: value, Next: next}
}

func (l *List) Length() int {
	length := 1
	cursor := l
	for cursor != nil {
		length++
		cursor = cursor.Next
	}
	return length
}

func (l *List) Display() string {
	values := []string{}
	cursor := l
	for cursor != nil {
		values = append(values, fmt.Sprintf("%d", cursor.Value))
		cursor = cursor.Next
	}
	return strings.Join(values, "->")
}

func (l *List) HasEvenLength() bool {
	cursor := l
	for cursor != nil && cursor.Next != nil {
		cursor = cursor.Next.Next
	}
	return cursor == nil
}

func (l *List) KthElementFromEnd(k uint) (int, bool) {
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

func (l *List) HasLoop() bool {
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
Why does it work? (for lists with loop)
- lets assume the distance from head to loop beginning is n
- lets assume the length of the loop is L

- when the slow pointer moves n, the faster pointer will have moved 2n
- if we consider the loop beginning to be at position 0:
  - the slow pointer will be at position 0
  - faster pointer will be at position n
  - for the faster pointer to reach the loop beginning again, it must move (L - n)
  - we want to find x, where 0 + x <=> (L - n) + 2x
  - so x = L - n + 2x =>
  - x + n = L + 2x =>
  - n = L + x
  - since L is a complete loop, safe to assume n = x
  - this means slow and fast pointers will meet at position n

- moving slow back to the head (position = -n)
  - fast pointer will be at position n => (L - n) from the loop beginning
  - this time, moving both pointers 1 by 1, after n steps:
  - slow pointer will be at position 0
  - fast pointer will be at position (L - n) + n => L <=> 0
*/
func (l *List) LoopStart() *List {
	slowPtr := l
	fastPtr := l
	hasLoop := false

	for fastPtr.Next != nil && fastPtr.Next.Next != nil {
		slowPtr = slowPtr.Next
		fastPtr = fastPtr.Next.Next
		if slowPtr == fastPtr {
			hasLoop = true
			break
		}
	}

	if !hasLoop {
		return nil
	}

	slowPtr = l
	for slowPtr != fastPtr {
		slowPtr = slowPtr.Next
		fastPtr = fastPtr.Next
	}

	return slowPtr
}

func (l *List) LoopLength() int {
	slowPtr := l
	fastPtr := l
	hasLoop := false

	for fastPtr.Next != nil && fastPtr.Next.Next != nil {
		slowPtr = slowPtr.Next
		fastPtr = fastPtr.Next.Next
		if slowPtr == fastPtr {
			hasLoop = true
			break
		}
	}

	if !hasLoop {
		return 0
	}

	length := 1
	for fastPtr.Next != slowPtr {
		fastPtr = fastPtr.Next
		length++
	}
	return length
}

func (l *List) Middle() int {
	slowPtr := l
	fastPtr := l
	for fastPtr != nil && fastPtr.Next != nil {
		slowPtr = slowPtr.Next
		fastPtr = fastPtr.Next.Next
	}
	return slowPtr.Value
}

func (l *List) IsPalindrome() bool {
	if l.Next == nil {
		return true
	}

	slowPtr := l
	fastPtr := l
	for fastPtr.Next != nil && fastPtr.Next.Next != nil {
		slowPtr = slowPtr.Next
		fastPtr = fastPtr.Next.Next
	} //O(n/2)

	rightHalf := ListReverse(slowPtr.Next) //O(n/2)
	slowPtr.Next = nil

	leftPtr := l
	rightPtr := rightHalf
	for rightPtr != nil {
		if leftPtr.Value != rightPtr.Value {
			slowPtr.Next = ListReverse(rightHalf)
			return false
		}
		leftPtr = leftPtr.Next
		rightPtr = rightPtr.Next
	} //O(n/2)

	slowPtr.Next = ListReverse(rightHalf) //O(n/2)
	return true
} //O(4n/2) -> O(2n) -> O(n)

func (l *List) ReverseDisplay() string {
	if l.Next == nil {
		return fmt.Sprintf("%d", l.Value)
	}

	return l.Next.ReverseDisplay() + fmt.Sprintf("->%d", l.Value)
}

func ListAppend(l *List, r *List) *List {
	if l == nil {
		return r
	}

	cursor := l
	for cursor.Next != nil {
		cursor = cursor.Next
	}
	cursor.Next = r
	return l
}

func ListMoveEvensLeft(l *List) *List {
	if l == nil {
		return nil
	}

	var evenList *List
	var oddList *List
	cursor := l
	for cursor != nil {
		bkp := cursor.Next
		cursor.Next = nil
		if cursor.Value%2 == 0 {
			evenList = ListAppend(evenList, cursor)
		} else {
			oddList = ListAppend(oddList, cursor)
		}
		cursor = bkp
	}
	return ListAppend(evenList, oddList)
}

func ListsFindMergePoint(l1 *List, l2 *List) *List {
	if l1 == nil || l2 == nil {
		return nil
	}

	l1Current := l1
	l2Current := l2

	l1Len := l1.Length() //O(n)
	l2Len := l2.Length() //O(n)

	if l1Len > l2Len {
		diff := l1Len - l2Len
		for range diff {
			l1Current = l1Current.Next //O(diff) => O(n)
		}
	} else {
		diff := l2Len - l1Len
		for range diff {
			l2Current = l2Current.Next //O(diff) => O(n)
		}
	}

	for l1Current != nil && l2Current != nil {
		if l1Current == l2Current {
			return l1Current
		}
		l1Current = l1Current.Next
		l2Current = l2Current.Next
	} //O(n)

	return nil
} //O(3n) => O(n)

func ListsMergeSorted(l1 *List, l2 *List) *List {
	// for simplicity, I will assume that both lists are sorted

	if l1 == nil {
		return l2
	}

	if l2 == nil {
		return l1
	}

	if l1.Value < l2.Value {
		l1.Next = ListsMergeSorted(l1.Next, l2)
		return l1
	}

	l2.Next = ListsMergeSorted(l1, l2.Next)
	return l2
}

func ListReverse(l *List) *List {
	if l == nil {
		return nil
	}

	var newHead *List
	ptr := l
	for ptr != nil {
		bkp := ptr.Next
		ptr.Next = newHead
		newHead = ptr
		ptr = bkp
	}
	return newHead
}

func ListReverseInPairs(l *List) *List {
	if l == nil || l.Next == nil {
		return l
	}

	bkp := l.Next.Next
	newHead := l.Next
	newHead.Next = l
	l.Next = ListReverseInPairs(bkp)

	return newHead
}

func ListReverseInGroups(l *List, k int) *List {
	// made a decision to not test validation errors like this one
	if k < 1 {
		panic("k must be greater than 0")
	}

	if l == nil || l.Next == nil || k == 1 {
		return l
	}

	count := 1
	newHead := l
	cursor := l.Next
	for cursor != nil && count < k {
		count++
		bkp := cursor.Next
		cursor.Next = newHead
		newHead = cursor
		l.Next = bkp
		cursor = bkp
	}

	l.Next = ListReverseInGroups(cursor, k)
	return newHead
}

func ListDeletePointer(l *List) {
	// made a decision to not test validation errors
	if l.Next == nil {
		panic("pointer cannot be at the end of the list")
	}

	l.Value = l.Next.Value
	l.Next = l.Next.Next
}

func SplitAndReverseMerge(l *List) *List {
	if l == nil || l.Next == nil {
		return l
	}

	slowPtr := l
	fastPtr := l.Next
	for fastPtr != nil && fastPtr.Next != nil {
		slowPtr = slowPtr.Next
		fastPtr = fastPtr.Next.Next
	}

	rightHalf := ListReverse(slowPtr.Next)
	slowPtr.Next = nil

	cursor := l
	for cursor != nil && rightHalf != nil {
		nextLeft := cursor.Next
		nextRight := rightHalf.Next

		cursor.Next = rightHalf
		rightHalf.Next = nextLeft

		rightHalf = nextRight
		cursor = nextLeft
	}

	return l
}

func PartitionList(l *List, k int) *List {
	if l == nil {
		return nil
	}

	var smallerList *List
	var greaterList *List
	cursor := l
	for cursor != nil {
		bkp := cursor.Next
		cursor.Next = nil
		if cursor.Value < k {
			smallerList = ListAppend(smallerList, cursor)
		} else {
			greaterList = ListAppend(greaterList, cursor)
		}
		cursor = bkp
	}
	return ListAppend(smallerList, greaterList)
}

func RemoveDuplicates(l *List) *List {
	if l == nil {
		return nil
	}

	prev := l
	cursor := l.Next
	for cursor != nil {
		if prev.Value == cursor.Value {
			prev.Next = cursor.Next
		} else {
			prev = prev.Next
		}
		cursor = cursor.Next
	}

	return l
}
