package select_sort

// Sort 选择排序是不稳定排序
func Sort(nums []int) {
	for i := 0; i < len(nums); i++ {
		minIdx := i
		for j := i + 1; j < len(nums); j++ {
			if nums[minIdx] > nums[j] {
				minIdx = j
			}
		}

		nums[i], nums[minIdx] = nums[minIdx], nums[i]
	}
}
