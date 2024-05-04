package heap

import (
	"github.com/stretchr/testify/require"
	"testing"
)

func TestMinHeap_Peek(t *testing.T) {
	heap := NewMinHeap(3)

	r, ok := heap.Peek()
	require.False(t, ok)
	require.Equal(t, -1, r)

	heap.Add(1)

	r, ok = heap.Peek()

	require.True(t, ok)
	require.Equal(t, 1, r)

}

func TestMinHeap_Size(t *testing.T) {
	heap := NewMinHeap(3)

	require.Equal(t, 0, heap.Size())

	heap.Add(2)
	require.Equal(t, 1, heap.Size())

	heap.Add(1)
	require.Equal(t, 2, heap.Size())
}

func TestMinHeap_Add(t *testing.T) {
	heap := NewMinHeap(3)
	heap.Add(2)
	heap.Add(3)
	heap.Add(1)

	r, ok := heap.Peek()

	require.True(t, ok)
	require.Equal(t, 1, r)
}

func TestMinHeap_Add2(t *testing.T) {
	heap := NewMinHeap(3)
	heap.Add(2)
	heap.Add(3)
	heap.Add(4)
	heap.Add(5)
	heap.Add(1)

	r, ok := heap.Peek()

	require.True(t, ok)
	require.Equal(t, 1, r)
}

func TestMinHeap_Remove(t *testing.T) {
	heap := NewMinHeap(3)
	heap.Add(2)
	heap.Add(3)
	heap.Add(4)
	heap.Add(5)
	heap.Add(1)

	require.Equal(t, 5, heap.Size())

	r, ok := heap.Remove()
	require.True(t, ok)
	require.Equal(t, 1, r)

	require.Equal(t, 4, heap.Size())
}
