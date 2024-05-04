package tree

import (
	"github.com/stretchr/testify/require"
	"testing"
	"util"
)

func TestTrie(t *testing.T) {
	trie := NewTrie()

	trie.Insert("我是中国人")
	trie.Insert("hello, world")
	trie.Insert("hi")

	require.Equal(t, true, trie.Search("我是中国人"))
	require.Equal(t, false, trie.Search("我是中国"))

	require.Equal(t, false, trie.Search("hello"))
	require.Equal(t, true, trie.StartWith("hello"))
	require.Equal(t, true, trie.Search("hello, world"))

	candidates := trie.Candidates("h")
	require.Equal(t, 2, len(candidates))
	require.Equal(t, true, util.Contains(candidates, 'i'))
	require.Equal(t, true, util.Contains(candidates, 'e'))

}
