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

func StackReverse(s *Stack) *Stack {
	// Time Complexity: O(n) | Space Complexity: O(n)
	// we could get space complexity O(1) by only reverting the direction of pointers
	// but the challenge was to only use push and pop

	// not using push and pop, but another idea on how to solve the problem
	// if we implemented the stack with a dynamic array,
	// we could have two pointers: leftPtr and a rightPtr
	// leftPtr starts at 0
	// rightPtr starts at len(array) - 1
	// leftPtr moves from left to right
	// rightPtr moves from right to left
	// swap the values from leftPtr and rightPtr and move them until rightPtr >= leftPtr
	// Time Complexity: O(n/2) => O(n) | Space Complexity: O(1)

	var newStack *Stack
	var poppedValue int
	for s != nil {
		s, poppedValue = StackPop(s)
		newStack = StackPush(newStack, poppedValue)
	}

	return newStack
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
