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
