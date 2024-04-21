package array

import (
	"github.com/stretchr/testify/require"
	"testing"
)

func TestBinaryIndexedTree_SumRange(t *testing.T) {
	nums := []int{8, 4, 6, 7, 9, 3, 1, 5}
	biTree := NewBinaryIndexedTree(nums)

	sum, _ := biTree.SumRange(1, 5)
	require.Equal(t, sum, 29)
}

func TestBinaryIndexedTree_Update(t *testing.T) {
	nums := []int{8, 4, 6, 7, 9, 3, 1, 5}
	biTree := NewBinaryIndexedTree(nums)

	_ = biTree.Update(7, 10)

	sum, _ := biTree.SumRange(1, 5)
	require.Equal(t, sum, 29)

	_ = biTree.Update(3, 10)
	sum, _ = biTree.SumRange(1, 5)
	require.Equal(t, sum, 32)
}
