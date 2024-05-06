package bit_operation

// SingleNumber3 https://leetcode.com/problems/single-number-iii
func SingleNumber3(nums []int) (r1, r2 int) {
	xor := 0
	for _, num := range nums {
		xor ^= num
	}

	lsb := xor & -xor
	r1, r2 = 0, 0

	for _, num := range nums {
		if num&lsb > 0 {
			r1 ^= num
		} else {
			r2 ^= num
		}
	}

	return
}
