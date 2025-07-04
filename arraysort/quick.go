package arraysort

func QuickSort(a []int) {
	if len(a) == 0 || len(a) == 1 {
		return
	}

	pivot := a[len(a)-1]
	swapMarker := -1
	for i := range len(a) {
		if i == len(a)-1 {
			swapMarker++
			swap(a, swapMarker, i)
		}

		if a[i] < pivot {
			swapMarker++
			if swapMarker != i {
				swap(a, swapMarker, i)
			}
		}
	}

	QuickSort(a[:swapMarker])
	QuickSort(a[swapMarker+1:])
}
