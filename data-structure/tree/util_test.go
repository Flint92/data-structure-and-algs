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
