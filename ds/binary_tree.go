package ds

import "fmt"

type BinaryTree struct {
	value int
	left  *BinaryTree
	right *BinaryTree
}

func NewBinaryTree(value int, left *BinaryTree, right *BinaryTree) *BinaryTree {
	return &BinaryTree{value: value, left: left, right: right}
}

func (t *BinaryTree) DisplayPreorder() string {
	result := fmt.Sprintf("%d", t.value)
	if t.left != nil {
		result += " " + t.left.DisplayPreorder()
	}

	if t.right != nil {
		result += " " + t.right.DisplayPreorder()
	}

	return result
}

func (t *BinaryTree) DisplayInorder() string {
	result := ""
	if t.left != nil {
		result += t.left.DisplayInorder() + " "
	}

	result += fmt.Sprintf("%d", t.value)

	if t.right != nil {
		result += " " + t.right.DisplayInorder()
	}

	return result
}

func (t *BinaryTree) DisplayPostorder() string {
	result := ""
	if t.left != nil {
		result += t.left.DisplayPostorder() + " "
	}

	if t.right != nil {
		result += t.right.DisplayPostorder() + " "
	}

	result += fmt.Sprintf("%d", t.value)
	return result
}

func (t *BinaryTree) DisplayLevelorder() string {
	result := ""
	stack := []*BinaryTree{t}
	for len(stack) > 0 {
		if len(result) > 0 {
			result += " "
		}
		node := stack[0]
		result += fmt.Sprintf("%d", node.value)
		if node.left != nil {
			stack = append(stack, node.left)
		}
		if node.right != nil {
			stack = append(stack, node.right)
		}
		if len(stack) > 0 {
			stack = stack[1:]
		}
	}
	return result
}

func (t *BinaryTree) Max() int {
	maxValue := t.value
	if t.left != nil {
		leftMax := t.left.Max()
		if leftMax > maxValue {
			maxValue = leftMax
		}
	}
	if t.right != nil {
		rightMax := t.right.Max()
		if rightMax > maxValue {
			maxValue = rightMax
		}
	}
	return maxValue
}

func (t *BinaryTree) Contains(value int) bool {
	return t.value == value || (t.left != nil && t.left.Contains(value)) || (t.right != nil && t.right.Contains(value))
}
