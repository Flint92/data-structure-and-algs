package tree

import (
	"github.com/stretchr/testify/require"
	"testing"
)

func TestRBTree_Insert(t *testing.T) {
	rbTree := NewRBTree()
	rbTree.Insert(5)

	require.Equal(t, 5, rbTree.root.val)
	require.Equal(t, BLACK, rbTree.root.color)

	rbTree.Insert(7)
	require.Equal(t, 5, rbTree.root.val)
	require.Equal(t, BLACK, rbTree.root.color)
	require.Equal(t, 7, rbTree.root.right.val)
	require.Equal(t, RED, rbTree.root.right.color)

	rbTree.Insert(4)
	require.Equal(t, 5, rbTree.root.val)
	require.Equal(t, BLACK, rbTree.root.color)
	require.Equal(t, 4, rbTree.root.left.val)
	require.Equal(t, RED, rbTree.root.left.color)

	rbTree.Insert(20)
	require.Equal(t, 5, rbTree.root.val)
	require.Equal(t, BLACK, rbTree.root.color)
	require.Equal(t, 4, rbTree.root.left.val)
	require.Equal(t, BLACK, rbTree.root.left.color)
	require.Equal(t, 7, rbTree.root.right.val)
	require.Equal(t, BLACK, rbTree.root.right.color)
	require.Equal(t, RED, rbTree.root.right.right.color)
	require.Equal(t, 20, rbTree.root.right.right.val)

	rbTree.Insert(6)
	require.Equal(t, 5, rbTree.root.val)
	require.Equal(t, BLACK, rbTree.root.color)
	require.Equal(t, 4, rbTree.root.left.val)
	require.Equal(t, BLACK, rbTree.root.left.color)
	require.Equal(t, 7, rbTree.root.right.val)
	require.Equal(t, BLACK, rbTree.root.right.color)
	require.Equal(t, RED, rbTree.root.right.right.color)
	require.Equal(t, 20, rbTree.root.right.right.val)
	require.Equal(t, RED, rbTree.root.right.left.color)
	require.Equal(t, 6, rbTree.root.right.left.val)

	rbTree.Insert(8)
	rbTree.Insert(21)
	rbTree.Insert(10)

	require.Equal(t, 7, rbTree.root.val)
	require.Equal(t, BLACK, rbTree.root.color)

	require.Equal(t, 5, rbTree.root.left.val)
	require.Equal(t, RED, rbTree.root.left.color)
	require.Equal(t, 4, rbTree.root.left.left.val)
	require.Equal(t, BLACK, rbTree.root.left.left.color)
	require.Equal(t, 6, rbTree.root.left.right.val)
	require.Equal(t, BLACK, rbTree.root.right.right.color)

	require.Equal(t, 20, rbTree.root.right.val)
	require.Equal(t, RED, rbTree.root.right.color)
	require.Equal(t, 8, rbTree.root.right.left.val)
	require.Equal(t, BLACK, rbTree.root.right.left.color)
	require.Equal(t, 10, rbTree.root.right.left.right.val)
	require.Equal(t, RED, rbTree.root.right.left.right.color)

	require.Equal(t, 21, rbTree.root.right.right.val)
	require.Equal(t, BLACK, rbTree.root.right.right.color)
}

func TestRBTree_Search(t *testing.T) {
	rbTree := NewRBTree()
	rbTree.Insert(1)
	rbTree.Insert(2)
	rbTree.Insert(3)

	require.Equal(t, true, rbTree.Search(1))
	require.Equal(t, false, rbTree.Search(4))
}

func TestRBTree_DELETE(t *testing.T) {
	rbTree := NewRBTree()
	rbTree.Insert(2)
	rbTree.Insert(1)
	rbTree.Insert(3)
	rbTree.Insert(4)

	require.Equal(t, false, rbTree.Delete(5))
	require.Equal(t, true, rbTree.Delete(4))
	require.Equal(t, true, rbTree.Delete(2))

	require.Equal(t, 3, rbTree.root.val)
	require.Equal(t, BLACK, rbTree.root.color)

	require.Equal(t, 1, rbTree.root.left.val)
	require.Equal(t, RED, rbTree.root.left.color)
}
