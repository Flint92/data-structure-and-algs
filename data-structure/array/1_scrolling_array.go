package array

func Fibonacci(n uint) uint {
	if n == 0 || n == 1 {
		return 1
	}

	var nums [3]uint
	nums[0] = 1
	nums[1] = 1

	for i := 2; i <= int(n); i++ {
		// using scrolling array
		nums[i%3] = nums[(i-1)%3] + nums[(i-2)%3]
	}

	return nums[n%3]
}
