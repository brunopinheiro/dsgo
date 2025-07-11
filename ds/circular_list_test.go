package ds_test

import (
	"testing"

	"github.com/brunopinheiro/dsgo/ds"
	"github.com/stretchr/testify/require"
)

func TestCircularList(t *testing.T) {
	fromArray := func(values []int) *ds.CircularListNode {
		if len(values) == 0 {
			return nil
		}

		l := ds.NewCircularList(values[0])
		for _, v := range values[1:] {
			l.Append(v)
		}
		return l
	}

	t.Run("creates a circular list with single element", func(t *testing.T) {
		l := ds.NewCircularList(1)
		require.Equal(t, "1", l.Display())
	})

	t.Run("append", func(t *testing.T) {
		l := ds.NewCircularList(1)
		l.Append(2)
		l.Append(3)
		require.Equal(t, "1->2->3", l.Display())
	})

	t.Run("split circular list into two", func(t *testing.T) {
		t.Run("with empty list", func(t *testing.T) {
			left, right := ds.SplitCircularListInHalf(nil)
			require.Nil(t, left)
			require.Nil(t, right)
		})

		t.Run("with single element", func(t *testing.T) {
			l := fromArray([]int{1})
			left, right := ds.SplitCircularListInHalf(l)
			require.Equal(t, "1", left.Display())
			require.Nil(t, right)
		})

		t.Run("with odd length", func(t *testing.T) {
			l := fromArray([]int{1, 2, 3, 4, 5})
			left, right := ds.SplitCircularListInHalf(l)
			require.Equal(t, "1->2->3", left.Display())
			require.Equal(t, "4->5", right.Display())
		})

		t.Run("with even length", func(t *testing.T) {
			l := fromArray([]int{1, 2, 3, 4, 5, 6})
			left, right := ds.SplitCircularListInHalf(l)
			require.Equal(t, "1->2->3", left.Display())
			require.Equal(t, "4->5->6", right.Display())
		})
	})
}
