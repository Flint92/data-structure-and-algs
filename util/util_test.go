package util

import (
	"github.com/stretchr/testify/require"
	"testing"
)

func TestMaxInt(t *testing.T) {
	require.Equal(t, 5, MaxInt(3, 5))
}

func TestMinInt(t *testing.T) {
	require.Equal(t, 1, MinInt(1, 3))
}

func TestContains(t *testing.T) {
	s := []int{1, 2, 3, 4, 5}
	require.True(t, Contains(s, 3))
	require.False(t, Contains(s, 6))
}
