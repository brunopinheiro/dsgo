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

func (t *BinaryTree) hasLeft() bool {
	return t.left != nil
}

func (t *BinaryTree) hasRight() bool {
	return t.right != nil
}

func (t *BinaryTree) isLeaf() bool {
	return t.left == nil && t.right == nil
}

func (t *BinaryTree) DisplayPreOrder() string {
	// without recursion, we could solve this problem using a queue
	result := fmt.Sprintf("%d", t.value)
	if t.hasLeft() {
		result += " " + t.left.DisplayPreOrder()
	}

	if t.hasRight() {
		result += " " + t.right.DisplayPreOrder()
	}

	return result
}

func (t *BinaryTree) DisplayInOrder() string {
	// without recursion, we could solve this problem using a queue
	result := ""
	if t.hasLeft() {
		result += t.left.DisplayInOrder() + " "
	}

	result += fmt.Sprintf("%d", t.value)

	if t.hasRight() {
		result += " " + t.right.DisplayInOrder()
	}

	return result
}

func (t *BinaryTree) DisplayPostOrder() string {
	// without recursion, we could solve this problem using a queue
	result := ""
	if t.hasLeft() {
		result += t.left.DisplayPostOrder() + " "
	}

	if t.hasRight() {
		result += t.right.DisplayPostOrder() + " "
	}

	result += fmt.Sprintf("%d", t.value)
	return result
}

func (t *BinaryTree) DisplayLevelOrder() string {
	result := ""
	queue := []*BinaryTree{t}
	for len(queue) > 0 {
		if len(result) > 0 {
			result += " "
		}
		node := queue[0]
		result += fmt.Sprintf("%d", node.value)
		if node.hasLeft() {
			queue = append(queue, node.left)
		}
		if node.hasRight() {
			queue = append(queue, node.right)
		}
		if len(queue) > 0 {
			queue = queue[1:]
		}
	}
	return result
}

func (t *BinaryTree) DisplayReverseLevelOrder() string {
	queue := []*BinaryTree{t}
	stack := []*BinaryTree{t}
	for len(queue) > 0 {
		node := queue[0]
		if node.hasRight() {
			stack = append(stack, node.right)
			queue = append(queue, node.right)
		}
		if node.hasLeft() {
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
	if t.hasLeft() {
		leftMax := t.left.Max()
		if leftMax > maxValue {
			maxValue = leftMax
		}
	}
	if t.hasRight() {
		rightMax := t.right.Max()
		if rightMax > maxValue {
			maxValue = rightMax
		}
	}
	return maxValue
}

func (t *BinaryTree) Contains(value int) bool {
	// without recursion, we could solve this problem using a queue
	return t.value == value || (t.hasLeft() && t.left.Contains(value)) || (t.hasRight() && t.right.Contains(value))
}

func (t *BinaryTree) Add(value int) {
	queue := []*BinaryTree{t}
	for {
		node := queue[0]
		queue = queue[1:]
		if node.isLeaf() {
			node.left = NewBinaryTree(value, nil, nil)
			break
		}
		if node.hasLeft() {
			if node.hasRight() {
				node.right = NewBinaryTree(value, nil, nil)
				break
			}
			queue = append(queue, node.left)
		}
		if node.hasRight() {
			queue = append(queue, node.right)
		}
	}
}

func (t *BinaryTree) Size() int {
	// without recursion, we could solve this problem using a queue
	size := 1
	if t.hasLeft() {
		size += t.left.Size()
	}
	if t.hasRight() {
		size += t.right.Size()
	}
	return size
}

func (t *BinaryTree) Height() int {
	leftHeight := 0
	if t.hasLeft() {
		leftHeight = t.left.Height()
	}

	rightHeight := 0
	if t.hasRight() {
		rightHeight = t.right.Height()
	}

	return 1 + max(leftHeight, rightHeight)
}

func (t *BinaryTree) DeepestValue() int {
	queue := []*BinaryTree{t}
	deepestNode := t
	for len(queue) > 0 {
		deepestNode = queue[0]
		queue = queue[1:]
		if deepestNode.hasLeft() {
			queue = append(queue, deepestNode.left)
		}
		if deepestNode.hasRight() {
			queue = append(queue, deepestNode.right)
		}
	}
	return deepestNode.value
}

func (t *BinaryTree) LeafCount() int {
	queue := []*BinaryTree{t}
	leafs := 0
	for len(queue) > 0 {
		node := queue[0]
		queue = queue[1:]
		if node.isLeaf() {
			leafs += 1
			continue
		}
		if node.hasLeft() {
			queue = append(queue, node.left)
		}
		if node.hasRight() {
			queue = append(queue, node.right)
		}
	}
	return leafs
}

func (t *BinaryTree) MaxLevelSum() int {
	queue := []*BinaryTree{t}
	maxSum := 0
	sumQueue := func(queue []*BinaryTree) int {
		sum := 0
		for _, node := range queue {
			sum += node.value
		}
		return sum
	}
	for len(queue) > 0 {
		nextQueue := []*BinaryTree{}
		levelSum := sumQueue(queue)
		if levelSum > maxSum {
			maxSum = levelSum
		}
		for _, node := range queue {
			if node.hasLeft() {
				nextQueue = append(nextQueue, node.left)
			}
			if node.hasRight() {
				nextQueue = append(nextQueue, node.right)
			}
		}
		queue = nextQueue
	}
	return maxSum
}

func (t *BinaryTree) RootToLeafPaths() [][]int {
	return binaryTreeRootToLeafPath(t, []*BinaryTree{})
}

func binaryTreeRootToLeafPath(t *BinaryTree, path []*BinaryTree) [][]int {
	arrayFromStack := func(stack []*BinaryTree) []int {
		result := []int{}
		for _, node := range stack {
			result = append(result, node.value)
		}
		return result
	}

	if t == nil {
		return [][]int{}
	}

	newPath := append(path, t)
	if t.isLeaf() {
		return [][]int{arrayFromStack(newPath)}
	}

	paths := [][]int{}
	if t.hasLeft() {
		paths = append(paths, binaryTreeRootToLeafPath(t.left, newPath)...)
	}
	if t.hasRight() {
		paths = append(paths, binaryTreeRootToLeafPath(t.right, newPath)...)
	}
	return paths
}

func (t *BinaryTree) HasPathWithSum(n int) bool {
	if t.value == n {
		return true
	}

	if t.value > n || t.isLeaf() {
		return false
	}

	if t.hasLeft() {
		if t.left.HasPathWithSum(n - t.value) {
			return true
		}
	}

	return t.hasRight() && t.right.HasPathWithSum(n-t.value)
}

func (t *BinaryTree) Sum() int {
	leftSum := 0
	if t.hasLeft() {
		leftSum = t.left.Sum()
	}

	rightSum := 0
	if t.hasRight() {
		rightSum = t.right.Sum()
	}

	return t.value + leftSum + rightSum
}

func BinaryTreeEqual(left *BinaryTree, right *BinaryTree) bool {
	if left == nil && right == nil {
		return true
	}
	if left == nil || right == nil {
		return false
	}
	return left.value == right.value &&
		BinaryTreeEqual(left.left, right.left) &&
		BinaryTreeEqual(left.right, right.right)
}
