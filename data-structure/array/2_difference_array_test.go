package array

import (
	"github.com/stretchr/testify/require"
	"testing"
)

func TestDifferenceArray_Increment(t *testing.T) {
	nums := []int{8, 4, 6, 7, 9, 3, 1, 5}
	da := MakeDifferenceArray(nums)
	_, _ = da.Increment(3, 5, 4)
	result := da.Result()

	require.Equal(t, result, []int{8, 4, 6, 11, 13, 7, 1, 5})
}
