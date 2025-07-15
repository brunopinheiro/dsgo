package ds

import (
	"fmt"
	"strings"
)

type Stack struct {
	value int
	next  *Stack
}

func NewStack(value int) *Stack {
	return &Stack{value: value, next: nil}
}

func (s *Stack) Display() string {
	cursor := s
	values := []string{}
	for cursor != nil {
		values = append(values, fmt.Sprintf("%d", cursor.value))
		cursor = cursor.next
	}
	return strings.Join(values, "->")
}

func (s *Stack) Peek() int {
	return s.value
}

func StackPush(s *Stack, value int) *Stack {
	newStack := NewStack(value)
	newStack.next = s
	return newStack
}

func StackPop(s *Stack) (*Stack, int) {
	if s == nil {
		return nil, 0
	}
	return s.next, s.value
}

// this functions is not related to the Stack structure.
// why keep it here?
// implemented this function as a challenge from Stacks section
// so, keeping it here for now
func IsArrayPalindrome(values []int) bool {
	if len(values) <= 1 {
		return true
	}

	var s *Stack
	middleIndex := 0
	for i, v := range values {
		if v == 0 {
			middleIndex = i
			break
		}
		s = StackPush(s, v)
	}

	poppedValue := 0
	for i := middleIndex + 1; i < len(values); i++ {
		if s == nil {
			return false
		}

		s, poppedValue = StackPop(s)
		if poppedValue != values[i] {
			return false
		}
	}

	return s == nil
}
