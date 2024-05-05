package shell_sort

// ShellSort 希尔排序是一种不稳定的排序
func ShellSort(nums []int) {
	n := len(nums)
	h := 1
	for h < n/3 {
		h = 3*h + 1 // 1, 4, 13, 40, 121, 364, 1093...
	}

	for h >= 1 {
		for i := h; i < n; i++ {
			for j := i; j >= h && nums[j] < nums[j-h]; j -= h {
				nums[j], nums[j-h] = nums[j-h], nums[j]
			}
		}
		h /= 3
	}
}
