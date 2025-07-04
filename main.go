package main

import (
	"fmt"

	"github.com/brunopinheiro/dsgo/arraysort"
)

func main() {
	a := []int{5, 4, 3, 2, 1}
	arraysort.BubbleSort(a)
	fmt.Println(a)

	b := []int{6, 5, 4, 3, 2, 1}
	fmt.Println(arraysort.MergeSort(b))

	c := []int{3, 2, 5, 0, 1, 8, 7, 6, 9, 4}
	arraysort.QuickSort(c)
	fmt.Println(c)
}
