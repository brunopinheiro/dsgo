package ds_test

import (
	"testing"

	"github.com/brunopinheiro/dsgo/ds"
	"github.com/stretchr/testify/require"
)

func TestBinaryTree(t *testing.T) {
	t.Run("display preorder", func(t *testing.T) {
		bt := ds.NewBinaryTree(
			1,
			ds.NewBinaryTree(
				2,
				ds.NewBinaryTree(4, nil, nil),
				ds.NewBinaryTree(5, nil, nil),
			),
			ds.NewBinaryTree(
				3,
				ds.NewBinaryTree(6, nil, nil),
				ds.NewBinaryTree(7, nil, nil),
			),
		)

		require.Equal(t, "1 2 4 5 3 6 7", bt.DisplayPreorder())
	})

	t.Run("display in order", func(t *testing.T) {
		bt := ds.NewBinaryTree(
			1,
			ds.NewBinaryTree(
				2,
				ds.NewBinaryTree(4, nil, nil),
				ds.NewBinaryTree(5, nil, nil),
			),
			ds.NewBinaryTree(
				3,
				ds.NewBinaryTree(6, nil, nil),
				ds.NewBinaryTree(7, nil, nil),
			),
		)

		require.Equal(t, "4 2 5 1 6 3 7", bt.DisplayInorder())
	})

	t.Run("display postorder", func(t *testing.T) {
		bt := ds.NewBinaryTree(
			1,
			ds.NewBinaryTree(
				2,
				ds.NewBinaryTree(4, nil, nil),
				ds.NewBinaryTree(5, nil, nil),
			),
			ds.NewBinaryTree(
				3,
				ds.NewBinaryTree(6, nil, nil),
				ds.NewBinaryTree(7, nil, nil),
			),
		)

		require.Equal(t, "4 5 2 6 7 3 1", bt.DisplayPostorder())
	})
}
