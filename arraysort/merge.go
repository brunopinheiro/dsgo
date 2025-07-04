package arraysort

func MergeSort(a []int) []int {
	if len(a) == 0 || len(a) == 1 {
		return a
	}

	half := int(len(a) / 2)
	left := MergeSort(a[0:half])
	right := MergeSort(a[half:])

	b := []int{}
	for len(left) > 0 || len(right) > 0 {
		if len(left) == 0 {
			b = append(b, right[0])
			right = right[1:]
			continue
		}

		if len(right) == 0 {
			b = append(b, left[0])
			left = left[1:]
			continue
		}

		if left[0] < right[0] {
			b = append(b, left[0])
			left = left[1:]
		} else {
			b = append(b, right[0])
			right = right[1:]
		}
	}

	return b
}
