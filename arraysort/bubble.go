package arraysort

func BubbleSort(a []int) {
	if len(a) == 0 || len(a) == 1 {
		return
	}

	for i := len(a) - 1; i > 0; i-- {
		changed := false
		for j := range i {
			if a[j] > a[j+1] {
				swap(a, j, j+1)
				changed = true
			}
		}
		if !changed {
			break
		}
	}
}
