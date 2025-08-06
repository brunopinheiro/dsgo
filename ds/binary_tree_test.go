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

	t.Run("display levelorder (bft)", func(t *testing.T) {
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

		require.Equal(t, "1 2 3 4 5 6 7", bt.DisplayLevelorder())
	})

	t.Run("max element", func(t *testing.T) {
		bt := ds.NewBinaryTree(
			6,
			ds.NewBinaryTree(
				7,
				ds.NewBinaryTree(3, nil, nil),
				ds.NewBinaryTree(4, nil, nil),
			),
			ds.NewBinaryTree(
				3,
				ds.NewBinaryTree(5, nil, nil),
				ds.NewBinaryTree(2, nil, nil),
			),
		)

		require.Equal(t, 7, bt.Max())
	})

	t.Run("contains", func(t *testing.T) {
		bt := ds.NewBinaryTree(
			6,
			ds.NewBinaryTree(
				7,
				ds.NewBinaryTree(3, nil, nil),
				ds.NewBinaryTree(4, nil, nil),
			),
			ds.NewBinaryTree(
				3,
				ds.NewBinaryTree(5, nil, nil),
				ds.NewBinaryTree(2, nil, nil),
			),
		)

		require.False(t, bt.Contains(9))
		require.True(t, bt.Contains(3))
	})

	t.Run("add keeps the tree complete", func(t *testing.T) {
		bt := ds.NewBinaryTree(1, nil, nil)
		bt.Add(2)
		require.Equal(t, "1 2", bt.DisplayLevelorder())

		bt.Add(3)
		require.Equal(t, "1 2 3", bt.DisplayLevelorder())

		bt.Add(4)
		require.Equal(t, "1 2 3 4", bt.DisplayLevelorder())

		bt.Add(5)
		require.Equal(t, "1 2 3 4 5", bt.DisplayLevelorder())

		bt.Add(6)
		require.Equal(t, "1 2 3 4 5 6", bt.DisplayLevelorder())
	})

	t.Run("size", func(t *testing.T) {
		bt := ds.NewBinaryTree(1, nil, nil)
		require.Equal(t, 1, bt.Size())

		bt = ds.NewBinaryTree(
			6,
			ds.NewBinaryTree(
				7,
				ds.NewBinaryTree(3, nil, nil),
				ds.NewBinaryTree(4, nil, nil),
			),
			ds.NewBinaryTree(
				3,
				ds.NewBinaryTree(5, nil, nil),
				ds.NewBinaryTree(
					2,
					ds.NewBinaryTree(
						8,
						nil,
						ds.NewBinaryTree(9, nil, nil),
					),
					nil,
				),
			),
		)
		require.Equal(t, 9, bt.Size())
	})

	t.Run("display reverser level order", func(t *testing.T) {
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

		require.Equal(t, "4 5 6 7 2 3 1", bt.DisplayReverseLevelorder())
	})

	t.Run("height", func(t *testing.T) {
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
		require.Equal(t, 3, bt.Height())

		bt = ds.NewBinaryTree(
			1,
			ds.NewBinaryTree(
				2,
				ds.NewBinaryTree(
					4,
					ds.NewBinaryTree(
						5,
						ds.NewBinaryTree(6, nil, nil),
						nil,
					),
					nil,
				),
				nil,
			),
			ds.NewBinaryTree(
				3,
				ds.NewBinaryTree(7, nil, nil),
				nil,
			),
		)
		require.Equal(t, 5, bt.Height())

		bt = ds.NewBinaryTree(
			1,
			ds.NewBinaryTree(
				3,
				ds.NewBinaryTree(7, nil, nil),
				nil,
			),
			ds.NewBinaryTree(
				2,
				ds.NewBinaryTree(
					4,
					ds.NewBinaryTree(
						5,
						ds.NewBinaryTree(6, nil, nil),
						nil,
					),
					nil,
				),
				nil,
			),
		)
		require.Equal(t, 5, bt.Height())
	})
}
