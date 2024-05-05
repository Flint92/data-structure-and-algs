package binary_search

import (
	"github.com/stretchr/testify/require"
	"testing"
)

func TestBinarySearch(t *testing.T) {
	nums := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16, 17, 18, 19, 20}

	require.Equal(t, BinarySearch(nums, 1), 0)
	require.Equal(t, BinarySearch(nums, 10), 9)
	require.Equal(t, BinarySearch(nums, 20), 19)
	require.Equal(t, BinarySearch(nums, 21), -1)

}
