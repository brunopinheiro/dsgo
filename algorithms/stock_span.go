package algorithms

import (
	"github.com/brunopinheiro/dsgo/ds"
)

func StockSpan(values []int) []int {
	if len(values) == 0 {
		return []int{}
	}

	span := make([]int, len(values))
	var peakStack *ds.Stack
	for i := range len(values) {
		if i == 0 {
			span[i] = 1
			peakStack = ds.StackPush(peakStack, i)
			continue
		}

		for peakStack != nil && values[i] >= values[peakStack.Peek()] {
			peakStack, _ = ds.StackPop(peakStack)
		}

		previousIndex := -1
		if peakStack != nil {
			previousIndex = peakStack.Peek()
		}
		span[i] = i - previousIndex
		peakStack = ds.StackPush(peakStack, i)
	}

	return span
}
