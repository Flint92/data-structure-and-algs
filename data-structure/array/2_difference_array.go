package array

import (
	"errors"
)

// DifferenceArray https://zhuanlan.zhihu.com/p/635853214
type DifferenceArray struct {
	nums []int // 原数组
	diff []int // 差分数组
}

func MakeDifferenceArray(nums []int) *DifferenceArray {
	numsCopy := nums[:]
	diffCopy := make([]int, len(numsCopy))
	diffCopy[0] = numsCopy[0]

	for i := 1; i < len(numsCopy); i++ {
		diffCopy[i] = numsCopy[i] - numsCopy[i-1]
	}

	return &DifferenceArray{
		nums: numsCopy,
		diff: diffCopy,
	}
}

func (da *DifferenceArray) Increment(a uint, b uint, val int) (bool, error) {
	if a > b {
		return false, errors.New("a must less than to b")
	}

	if int(a) >= len(da.nums) {
		return false, errors.New("a is out of range")
	}

	if int(b) >= len(da.nums) {
		return false, errors.New("b is out of range")
	}

	if val == 0 {
		return false, nil
	}

	da.diff[a] += val
	if int(b)+1 < len(da.diff) {
		da.diff[b+1] -= val
	}

	return true, nil
}

func (da *DifferenceArray) Result() []int {
	da.nums[0] = da.diff[0]
	for i := 1; i < len(da.diff); i++ {
		da.nums[i] = da.diff[i] + da.nums[i-1]
	}
	return da.nums[:]
}
