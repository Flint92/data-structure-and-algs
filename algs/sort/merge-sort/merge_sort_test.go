package merge_sort

import (
	"github.com/stretchr/testify/require"
	"testing"
)

func TestMergeSort(t *testing.T) {
	nums := []int{5, 4, 1, 3, 2, 8, 6, 9, 7}
	MergeSort(nums)
	require.Equal(t, []int{1, 2, 3, 4, 5, 6, 7, 8, 9}, nums)
}
