package bit_operation

import (
	"github.com/stretchr/testify/require"
	"testing"
)

func TestSingleNumber(t *testing.T) {
	nums1 := []int{2, 2, 1}
	require.Equal(t, SingleNumber(nums1), 1)

	nums2 := []int{4, 1, 2, 1, 2}
	require.Equal(t, SingleNumber(nums2), 4)

	nums3 := []int{1}
	require.Equal(t, SingleNumber(nums3), 1)
}
