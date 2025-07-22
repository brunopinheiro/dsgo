package algorithms

func MaxNumberInSlidingWindows(values []int, w int) []int {
	// a heap structure might be a good option to solve this problem,
	// were we always remove the leftmost value of the sliding window from the heap,
	// and add the new rightmost value. haven't implemented one yet, so, skipping for now
	if w > len(values) {
		panic("sliding window size cannot be greater than values.length")
	}

	return []int{}
}
