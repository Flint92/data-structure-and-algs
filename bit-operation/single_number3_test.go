package bit_operation

import (
	"github.com/stretchr/testify/require"
	"testing"
)

func TestSingleNumber3(t *testing.T) {
	nums1 := []int{1, 2, 1, 3, 2, 5}
	r1, r2 := SingleNumber3(nums1)

	t.Log(r1, r2)
	if r1 == 3 {
		require.Equal(t, r2, 5)
	} else {
		require.Equal(t, r1, 5)
		require.Equal(t, r2, 3)
	}

	nums2 := []int{-1, 0}
	r1, r2 = SingleNumber3(nums2)
	t.Log(r1, r2)
	if r1 == -1 {
		require.Equal(t, r2, 0)
	} else {
		require.Equal(t, r1, 0)
		require.Equal(t, r2, -1)
	}

	nums3 := []int{0, 1}
	r1, r2 = SingleNumber3(nums3)
	t.Log(r1, r2)
	if r1 == 0 {
		require.Equal(t, r2, 1)
	} else {
		require.Equal(t, r1, 1)
		require.Equal(t, r2, 0)
	}
}
