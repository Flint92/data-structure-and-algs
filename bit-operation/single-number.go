package bit_operation

// SingleNumber https://leetcode.com/problems/single-number/description
func SingleNumber(nums []int) int {
	result := 0
	for _, num := range nums {
		result ^= num
	}
	return result
}