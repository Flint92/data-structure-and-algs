package list

import (
	"github.com/stretchr/testify/require"
	"testing"
)

func TestSkipList(t *testing.T) {
	skipList := NewSkipList()
	for i := 1; i <= 10; i++ {
		skipList.Add(i)
	}

	require.Equal(t, skipList.Search(5), true)
	require.Equal(t, skipList.Search(51), false)
	skipList.Remove(5)
	require.Equal(t, skipList.Search(5), false)
}
