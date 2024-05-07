package tree

import (
	"github.com/stretchr/testify/require"
	"testing"
)

func TestAVLTree(t *testing.T) {
	avlTree := NewAVLTree()

	avlTree.Insert(3)
	avlTree.Insert(4)
	avlTree.Insert(5)
	avlTree.Insert(8)
	avlTree.Insert(9)

	require.Equal(t, avlTree.root.val, 4)
	require.Equal(t, avlTree.root.left.val, 3)
	require.Equal(t, avlTree.root.right.val, 8)
	require.Equal(t, avlTree.root.right.left.val, 5)
	require.Equal(t, avlTree.root.right.right.val, 9)
	require.Equal(t, avlTree.root.height, 3)

	require.Equal(t, avlTree.Insert(3), false)

	require.Equal(t, avlTree.Search(9), true)
	require.Equal(t, avlTree.Search(6), false)

	require.Equal(t, avlTree.Delete(6), false)
	require.Equal(t, avlTree.Delete(8), true)
}

func TestAVLTree_PreOrderTraversal(t *testing.T) {
	avlTree := NewAVLTree()

	avlTree.Insert(2)
	avlTree.Insert(3)
	avlTree.Insert(4)
	avlTree.Insert(5)
	avlTree.Insert(8)
	avlTree.Insert(9)

	preOrders := avlTree.PreOrderTraversal()
	require.Equal(t, []int{5, 3, 2, 4, 8, 9}, preOrders)
}

func TestAVLTree_InOrderTraversal(t *testing.T) {
	avlTree := NewAVLTree()

	avlTree.Insert(2)
	avlTree.Insert(3)
	avlTree.Insert(4)
	avlTree.Insert(5)
	avlTree.Insert(8)
	avlTree.Insert(9)

	inOrders := avlTree.InOrderTraversal()
	require.Equal(t, []int{2, 3, 4, 5, 8, 9}, inOrders)
}

func TestAVLTree_PostOrderTraversal(t *testing.T) {
	avlTree := NewAVLTree()

	avlTree.Insert(2)
	avlTree.Insert(3)
	avlTree.Insert(4)
	avlTree.Insert(5)
	avlTree.Insert(8)
	avlTree.Insert(9)

	postOrders := avlTree.PostOrderTraversal()
	require.Equal(t, []int{2, 4, 3, 9, 8, 5}, postOrders)
}
