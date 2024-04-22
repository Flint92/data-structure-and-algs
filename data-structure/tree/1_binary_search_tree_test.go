package tree

import (
	"github.com/stretchr/testify/require"
	"testing"
)

func TestBST(t *testing.T) {
	bst := NewBST()
	bst.Print()

	bst.Insert(3)
	bst.Insert(8)
	bst.Insert(4)
	bst.Insert(9)
	bst.Insert(5)

	require.Equal(t, bst.Insert(3), false)

	require.Equal(t, bst.Search(9), true)
	require.Equal(t, bst.Search(6), false)

	require.Equal(t, bst.Delete(6), false)
	require.Equal(t, bst.Delete(8), true)

	bst.Print()
}
