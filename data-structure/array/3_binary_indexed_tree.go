package array

import (
	"errors"
)

// BinaryIndexedTree
// https://zhuanlan.zhihu.com/p/99167607
// https://www.baeldung.com/cs/fenwick-tree
// 单点更新，区间查询
type BinaryIndexedTree struct {
	a []int // 原始数组，有效坐标从0开始
	c []int // 树状数组，有效坐标从1开始
}

func NewBinaryIndexedTree(a []int) *BinaryIndexedTree {
	aCopy := a[:]
	cCopy := make([]int, len(a)+1)

	for i := 0; i < len(aCopy); i++ {
		add(uint(i+1), a[i], cCopy)
	}

	return &BinaryIndexedTree{
		a: aCopy,
		c: cCopy,
	}
}

func (biTree *BinaryIndexedTree) Update(i uint, val int) error {
	if int(i) >= len(biTree.a) {
		return errors.New("index out of range")
	}
	add(i+1, val-biTree.a[i], biTree.c)
	biTree.a[i] = val
	return nil
}

func (biTree *BinaryIndexedTree) SumRange(left uint, right uint) (int, error) {
	if left > right {
		return -1, errors.New("left must be less than right")
	}

	if int(left) >= len(biTree.a) {
		return -1, errors.New("left is out of range")
	}

	if int(right) >= len(biTree.a) {
		return -1, errors.New("right is out of range")
	}

	if left == right {
		return 0, nil
	}

	return prefixSum(right+1, biTree.c) - prefixSum(left, biTree.c), nil
}

// 计算 [0, i-1]	区间的和
func prefixSum(n uint, c []int) int {
	var sum int

	for n > 0 {
		sum += c[n]
		n -= lowBit(n)
	}

	return sum
}

// c[i] = a[i - 2^k+1] + a[i - 2^k+2] + ... + a[i]
// k 为 i 的二进制中最右边连续 0 的个数
func add(i uint, val int, c []int) {
	for int(i) < len(c) {
		c[int(i)] += val
		i += lowBit(i)
	}
}

// 非负整数在二进制下最低位1以及后面的0构成的数值
func lowBit(n uint) uint {
	return n & (-n)
}
