package list

import (
	"github.com/stretchr/testify/require"
	"testing"
)

func TestXORList(t *testing.T) {
	xorList := NewXORList()
	xorList.Insert(1)
	xorList.Insert(2)
	xorList.Insert(3)
	xorList.Insert(4)
	xorList.Insert(5)

	require.Equal(t, xorList.Search(2), true)
	require.Equal(t, xorList.Search(6), false)

	xorList.Remove(3)
	require.Equal(t, xorList.Search(3), false)

	xorList.Remove(1)
	require.Equal(t, xorList.Search(1), false)

	xorList.Remove(5)
	require.Equal(t, xorList.Search(5), false)

	require.Equal(t, xorList.Search(2), true)
	require.Equal(t, xorList.Search(4), true)

	xorList.Remove(2)
	xorList.Remove(4)

	require.Equal(t, xorList.Search(2), false)
	require.Equal(t, xorList.Search(4), false)

}
