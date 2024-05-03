package tree

import (
	"github.com/stretchr/testify/require"
	"testing"
)

func TestMaxInt(t *testing.T) {
	require.Equal(t, 5, maxInt(3, 5))
}

func TestMinInt(t *testing.T) {
	require.Equal(t, 1, minInt(1, 3))
}

func TestContains(t *testing.T) {
	s := []int{1, 2, 3, 4, 5}
	require.True(t, contains(s, 3))
	require.False(t, contains(s, 6))
}
