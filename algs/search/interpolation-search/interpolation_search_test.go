package interpolation_search

import (
	"github.com/stretchr/testify/require"
	"testing"
)

func TestInterpolationSearch(t *testing.T) {
	nums := []int{1, 2, 93, 94, 95, 96, 97, 98, 99, 100, 101, 102, 103, 104, 105, 106, 107, 108, 109, 110}

	require.Equal(t, InterpolationSearch(nums, 1), 0)
	require.Equal(t, InterpolationSearch(nums, 100), 9)
	require.Equal(t, InterpolationSearch(nums, 0), -1)
	require.Equal(t, InterpolationSearch(nums, 200), -1)
	require.Equal(t, InterpolationSearch(nums, 20), -1)

}
