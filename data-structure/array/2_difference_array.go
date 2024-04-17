package main

import (
	"errors"
	"fmt"
)

func main() {
	nums := []int{8, 4, 6, 7, 9, 3, 1, 5}
	da := MakeDifferenceArray(nums)
	_, _ = da.Increment(3, 5, 4)
	result := da.Result()
	fmt.Println(result)
}

// DifferenceArray 差值数组
/*

- 差值数组：每个位置的值是原数组当前位置的值与前一个位置的值的差值，常用于对于某个区间内的所有值进行加减操作
- 以下是一个具体的例子：

| 下标 		| 0 | 1 | 2 | 3 | 4 | 5 | 6 | 7 |
| 原数组 	| 8 | 4 | 6 | 7 | 9 | 3 | 1 | 5 |
| 差分数组 	| 8 | -4| 2 | 1 | 2 | -6| -2| 4 |

- 可以看到 nums[0] = d[0], nums[3]=d[0]+d[1}+d[2]+d[3]
- 对区间 [a,b] 每个元素加减 ，不需要一个个操作，只需要在两端修改即可

*/
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
