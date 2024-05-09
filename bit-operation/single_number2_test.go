package bit_operation

import (
	"github.com/stretchr/testify/require"
	"testing"
)

func TestSingleNumber2(t *testing.T) {
	nums1 := []int{2, 2, 3, 2}
	require.Equal(t, SingleNumber2(nums1), 3)

	nums2 := []int{0, 1, 0, 1, 0, 1, 99}
	require.Equal(t, SingleNumber2(nums2), 99)

	nums3 := []int{0, 1, 0, 1, 0, 1, 1 << 31}
	require.Equal(t, SingleNumber2(nums3), -2147483648)

	nums4 := []int{0, 1, 0, 1, 0, 1, 1<<31 - 1}
	require.Equal(t, SingleNumber2(nums4), 2147483647)
}
