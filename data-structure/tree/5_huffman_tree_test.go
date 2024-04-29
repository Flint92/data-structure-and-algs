package tree

import (
	"github.com/stretchr/testify/require"
	"testing"
)

func TestHuffmanTree_Encode(t *testing.T) {
	hTree := NewHuffmanTree(map[string]float64{
		"a": 6.0,
		"b": 4.0,
		"c": 3.0,
		"d": 3.0,
	})

	source := "aabcddcaaaadcbbb"

	encodeStr := hTree.Encode(source)
	decodeStr, err := hTree.Decode(encodeStr)
	require.Nil(t, err)
	require.Equal(t, decodeStr, source)

	_, err = hTree.Decode("abc")
	require.NotNil(t, err)
	require.Equal(t, err.Error(), "unknown char a")

	_, err = hTree.Decode("110")
	require.NotNil(t, err)
	require.Equal(t, err.Error(), "invalid encode str 110")
}
