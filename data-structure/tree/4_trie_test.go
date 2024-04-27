package tree

import (
	"github.com/stretchr/testify/require"
	"testing"
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

	require.Equal(t, []rune{'e', 'i'}, trie.Candidates("h"))

}
