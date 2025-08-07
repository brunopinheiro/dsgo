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

	t.Run("deepest value", func(t *testing.T) {
		t.Run("when tree is full", func(t *testing.T) {
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
			require.Equal(t, 7, bt.DeepestValue())
		})

		t.Run("when tree is complete", func(t *testing.T) {
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
					nil,
				),
			)
			require.Equal(t, 6, bt.DeepestValue())
		})

		t.Run("when tree is unbalanced with left having deepest value", func(t *testing.T) {
			bt := ds.NewBinaryTree(
				1,
				ds.NewBinaryTree(
					2,
					ds.NewBinaryTree(4, nil, nil),
					ds.NewBinaryTree(
						5,
						ds.NewBinaryTree(
							7,
							nil,
							ds.NewBinaryTree(8, nil, nil),
						),
						nil,
					),
				),
				ds.NewBinaryTree(
					3,
					ds.NewBinaryTree(6, nil, nil),
					nil,
				),
			)
			require.Equal(t, 8, bt.DeepestValue())
		})

		t.Run("when tree is unbalanced with right having deepest value", func(t *testing.T) {
			bt := ds.NewBinaryTree(
				1,
				ds.NewBinaryTree(
					3,
					ds.NewBinaryTree(6, nil, nil),
					nil,
				),
				ds.NewBinaryTree(
					2,
					ds.NewBinaryTree(4, nil, nil),
					ds.NewBinaryTree(
						5,
						ds.NewBinaryTree(
							7,
							ds.NewBinaryTree(8, nil, nil),
							nil,
						),
						nil,
					),
				),
			)
			require.Equal(t, 8, bt.DeepestValue())
		})
	})

	t.Run("count number of leafs", func(t *testing.T) {
		bt := ds.NewBinaryTree(
			1,
			ds.NewBinaryTree(
				3,
				ds.NewBinaryTree(6, nil, nil),
				nil,
			),
			ds.NewBinaryTree(
				2,
				ds.NewBinaryTree(4, nil, nil),
				ds.NewBinaryTree(
					5,
					ds.NewBinaryTree(
						7,
						ds.NewBinaryTree(8, nil, nil),
						nil,
					),
					nil,
				),
			),
		)
		require.Equal(t, 3, bt.LeafCount())
	})

	t.Run("equal", func(t *testing.T) {
		t.Run("true when values and structures are equal", func(t *testing.T) {
			leftBT := ds.NewBinaryTree(
				1,
				ds.NewBinaryTree(
					3,
					ds.NewBinaryTree(6, nil, nil),
					nil,
				),
				ds.NewBinaryTree(
					2,
					ds.NewBinaryTree(4, nil, nil),
					ds.NewBinaryTree(
						5,
						ds.NewBinaryTree(
							7,
							ds.NewBinaryTree(8, nil, nil),
							nil,
						),
						nil,
					),
				),
			)

			rightBT := ds.NewBinaryTree(
				1,
				ds.NewBinaryTree(
					3,
					ds.NewBinaryTree(6, nil, nil),
					nil,
				),
				ds.NewBinaryTree(
					2,
					ds.NewBinaryTree(4, nil, nil),
					ds.NewBinaryTree(
						5,
						ds.NewBinaryTree(
							7,
							ds.NewBinaryTree(8, nil, nil),
							nil,
						),
						nil,
					),
				),
			)

			require.True(t, ds.BinaryTreeEqual(leftBT, rightBT))

		})

		t.Run("true when both are nil", func(t *testing.T) {
			require.True(t, ds.BinaryTreeEqual(nil, nil))
		})

		t.Run("false when structures are equal but values are different", func(t *testing.T) {
			leftBT := ds.NewBinaryTree(
				1,
				ds.NewBinaryTree(
					2,
					ds.NewBinaryTree(3, nil, nil),
					ds.NewBinaryTree(4, nil, nil),
				),
				ds.NewBinaryTree(
					5,
					ds.NewBinaryTree(6, nil, nil),
					ds.NewBinaryTree(7, nil, nil),
				),
			)

			rightBT := ds.NewBinaryTree(
				1,
				ds.NewBinaryTree(
					2,
					ds.NewBinaryTree(99, nil, nil),
					ds.NewBinaryTree(4, nil, nil),
				),
				ds.NewBinaryTree(
					5,
					ds.NewBinaryTree(6, nil, nil),
					ds.NewBinaryTree(7, nil, nil),
				),
			)

			require.False(t, ds.BinaryTreeEqual(leftBT, rightBT))
		})

		t.Run("false when values are equal but structures are different", func(t *testing.T) {
			leftBT := ds.NewBinaryTree(
				1,
				ds.NewBinaryTree(
					2,
					ds.NewBinaryTree(3, nil, nil),
					ds.NewBinaryTree(4, nil, nil),
				),
				ds.NewBinaryTree(
					5,
					ds.NewBinaryTree(6, nil, nil),
					ds.NewBinaryTree(7, nil, nil),
				),
			)

			rightBT := ds.NewBinaryTree(
				1,
				ds.NewBinaryTree(
					2,
					ds.NewBinaryTree(3, nil, nil),
					nil,
				),
				ds.NewBinaryTree(
					5,
					ds.NewBinaryTree(6, nil, nil),
					ds.NewBinaryTree(7, nil, nil),
				),
			)

			require.False(t, ds.BinaryTreeEqual(leftBT, rightBT))
		})

		t.Run("false when one of the trees is nil", func(t *testing.T) {
			require.False(t, ds.BinaryTreeEqual(nil, ds.NewBinaryTree(1, nil, nil)))
			require.False(t, ds.BinaryTreeEqual(ds.NewBinaryTree(1, nil, nil), nil))
		})
	})
}
