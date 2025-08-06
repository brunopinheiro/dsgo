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
	// without recursion, we could solve this problem using a queue
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
	// without recursion, we could solve this problem using a queue
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
	// without recursion, we could solve this problem using a queue
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
	queue := []*BinaryTree{t}
	for len(queue) > 0 {
		if len(result) > 0 {
			result += " "
		}
		node := queue[0]
		result += fmt.Sprintf("%d", node.value)
		if node.left != nil {
			queue = append(queue, node.left)
		}
		if node.right != nil {
			queue = append(queue, node.right)
		}
		if len(queue) > 0 {
			queue = queue[1:]
		}
	}
	return result
}

func (t *BinaryTree) DisplayReverseLevelorder() string {
	queue := []*BinaryTree{t}
	stack := []*BinaryTree{t}
	for len(queue) > 0 {
		node := queue[0]
		if node.right != nil {
			stack = append(stack, node.right)
			queue = append(queue, node.right)
		}
		if node.left != nil {
			stack = append(stack, node.left)
			queue = append(queue, node.left)
		}
		if len(queue) > 0 {
			queue = queue[1:]
		}
	}

	result := ""
	for len(stack) > 0 {
		top := stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		if len(result) > 0 {
			result += " "
		}
		result += fmt.Sprintf("%d", top.value)
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
	// without recursion, we could solve this problem using a queue
	return t.value == value || (t.left != nil && t.left.Contains(value)) || (t.right != nil && t.right.Contains(value))
}

func (t *BinaryTree) Add(value int) {
	queue := []*BinaryTree{t}
	for {
		node := queue[0]
		queue = queue[1:]
		if node.left == nil && node.right == nil {
			node.left = NewBinaryTree(value, nil, nil)
			break
		}
		if node.left != nil {
			if node.right == nil {
				node.right = NewBinaryTree(value, nil, nil)
				break
			}
			queue = append(queue, node.left)
		}
		if node.right != nil {
			queue = append(queue, node.right)
		}
	}
}

func (t *BinaryTree) Size() int {
	// without recursion, we could solve this problem using a queue
	size := 1
	if t.left != nil {
		size += t.left.Size()
	}
	if t.right != nil {
		size += t.right.Size()
	}
	return size
}

func (t *BinaryTree) Height() int {
	leftHeight := 0
	if t.left != nil {
		leftHeight = t.left.Height()
	}

	rightHeight := 0
	if t.right != nil {
		rightHeight = t.right.Height()
	}

	return 1 + max(leftHeight, rightHeight)
}
