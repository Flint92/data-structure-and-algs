package kmp

import (
	"github.com/stretchr/testify/require"
	"testing"
)

func TestKMP(t *testing.T) {
	s := "abcabdabeabc"
	p := "abc"

	r := KMP(s, p)
	require.Equal(t, []int{0, 9}, r)

	s = "世界新世界闻世"
	p = "世界"
	r = KMP(s, p)
	require.Equal(t, []int{0, 3}, r)
}

func TestNext(t *testing.T) {
	p := []rune("abcabea")
	t.Log(next(p))

	p = []rune("世界新世界闻世")
	t.Log(next(p))
}
