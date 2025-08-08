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

		require.Equal(t, "1 2 4 5 3 6 7", bt.DisplayPreOrder())
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

		require.Equal(t, "4 2 5 1 6 3 7", bt.DisplayInOrder())
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

		require.Equal(t, "4 5 2 6 7 3 1", bt.DisplayPostOrder())
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

		require.Equal(t, "1 2 3 4 5 6 7", bt.DisplayLevelOrder())
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
		require.Equal(t, "1 2", bt.DisplayLevelOrder())

		bt.Add(3)
		require.Equal(t, "1 2 3", bt.DisplayLevelOrder())

		bt.Add(4)
		require.Equal(t, "1 2 3 4", bt.DisplayLevelOrder())

		bt.Add(5)
		require.Equal(t, "1 2 3 4 5", bt.DisplayLevelOrder())

		bt.Add(6)
		require.Equal(t, "1 2 3 4 5 6", bt.DisplayLevelOrder())
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

		require.Equal(t, "4 5 6 7 2 3 1", bt.DisplayReverseLevelOrder())
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

	t.Run("max level sum", func(t *testing.T) {
		bt := ds.NewBinaryTree(
			1, // level 0
			ds.NewBinaryTree(
				3,                             // level 1
				ds.NewBinaryTree(6, nil, nil), // level 2
				nil,
			),
			ds.NewBinaryTree(
				2,                             // level 1
				ds.NewBinaryTree(4, nil, nil), // level 2
				ds.NewBinaryTree(
					5, // level 2
					ds.NewBinaryTree(
						7,                             // level 3
						ds.NewBinaryTree(8, nil, nil), // level 4
						nil,
					),
					nil,
				),
			),
		)

		require.Equal(t, 15, bt.MaxLevelSum()) // level 2
	})

	t.Run("all root to leaf paths", func(t *testing.T) {
		bt := ds.NewBinaryTree(
			1,
			ds.NewBinaryTree(
				3,
				ds.NewBinaryTree(6, nil, nil),
				nil,
			),
			ds.NewBinaryTree(
				2,
				ds.NewBinaryTree(
					4,
					ds.NewBinaryTree(11, nil, nil),
					ds.NewBinaryTree(9, nil, nil),
				),
				ds.NewBinaryTree(
					5,
					ds.NewBinaryTree(
						7,
						ds.NewBinaryTree(8, nil, nil),
						nil,
					),
					ds.NewBinaryTree(13, nil, nil),
				),
			),
		)

		require.ElementsMatch(t, [][]int{
			{1, 3, 6},
			{1, 2, 4, 9},
			{1, 2, 4, 11},
			{1, 2, 5, 7, 8},
			{1, 2, 5, 13},
		}, bt.RootToLeafPaths())
	})
}
