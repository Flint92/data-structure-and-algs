package tree

import (
	"github.com/stretchr/testify/require"
	"testing"
)

func TestSegmentTree_QueryRange(t *testing.T) {
	seg := NewSegmentTree([]int{1, 3, 5, 7, 9, 11})

	require.Equal(t, seg.QueryRange(0, 5), 36)
	require.Equal(t, seg.QueryRange(2, 2), 5)
	require.Equal(t, seg.QueryRange(2, 4), 21)

	seg.Print()
}

func TestSegmentTree_UpdateRange(t *testing.T) {
	seg := NewSegmentTree([]int{1, 3, 5, 7, 9, 11})
	seg.UpdateRange(2, 4, 1)

	require.Equal(t, seg.QueryRange(0, 5), 39)
	require.Equal(t, seg.QueryRange(2, 2), 6)
	require.Equal(t, seg.QueryRange(2, 4), 24)

	seg.Print()
}
