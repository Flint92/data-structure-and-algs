package quick_sort

// QuickSort 快排是一种不稳定排序
func QuickSort(nums []int) {
	quickSort(nums, 0, len(nums)-1)
}

func quickSort(nums []int, start, end int) {
	if start < end {
		j := partition(nums, start, end)
		quickSort(nums, start, j-1)
		quickSort(nums, j+1, end)
	}
}

func partition(nums []int, start int, end int) int {
	pivot := nums[start]
	for start < end {
		for start < end && nums[end] >= pivot {
			end--
		}
		nums[start], nums[end] = nums[end], nums[start]
		for start < end && nums[start] <= pivot {
			start++
		}
		nums[start], nums[end] = nums[end], nums[start]
	}
	return start
}
