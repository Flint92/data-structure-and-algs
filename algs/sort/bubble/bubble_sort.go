package bubble

func Sort(nums []int) {
	for i := 1; i < len(nums); i++ {
		ordered := true
		for j := 0; j < len(nums)-i; j++ {
			if nums[j] > nums[j+1] {
				nums[j], nums[j+1] = nums[j+1], nums[j]
				ordered = false
			}
		}
		if ordered {
			break
		}
	}
}
