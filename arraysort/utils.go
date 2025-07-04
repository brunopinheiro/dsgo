package arraysort

func swap(a []int, i, j int) {
	a[i] = a[i] ^ a[j]
	a[j] = a[i] ^ a[j]
	a[i] = a[i] ^ a[j]
}
