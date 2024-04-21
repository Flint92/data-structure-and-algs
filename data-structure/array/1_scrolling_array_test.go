package array

import (
	"github.com/stretchr/testify/require"
	"testing"
)

func TestFibonacci(t *testing.T) {
	f0 := Fibonacci(0)
	require.Equal(t, f0, uint(1))

	f1 := Fibonacci(1)
	require.Equal(t, f1, uint(1))

	for i := 2; i < 10; i++ {
		f := Fibonacci(uint(i))
		require.Equal(t, f, f0+f1)

		f0 = f1
		f1 = f
	}

}
