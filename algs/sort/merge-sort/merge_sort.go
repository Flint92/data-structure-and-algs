package merge_sort

func MergeSort(nums []int) {
	mergeSort(nums, 0, len(nums)-1)
}

func mergeSort(nums []int, start, end int) {
	if start < end {
		mid := start + (end-start)>>1
		mergeSort(nums, start, mid)
		mergeSort(nums, mid+1, end)
		merge(nums, start, mid, end)
	}
}

func merge(nums []int, start, mid, end int) {
	length := end - start + 1
	tmp := make([]int, length)
	index := 0
	start1, start2 := start, mid+1

	// 双指针合并两个有序数组
	for start1 <= mid && start2 <= end {
		if nums[start1] <= nums[start2] {
			tmp[index] = nums[start1]
			start1++
		} else {
			tmp[index] = nums[start2]
			start2++
		}
		index++
	}
	// 如果第一个数组后面还有数字就把他添加到临时数组中
	for start1 <= mid {
		tmp[index] = nums[start1]
		index++
		start1++
	}
	// 同理，如果第二个数组后面还有数字就把他添加到临时数组中
	for start2 <= end {
		tmp[index] = nums[start2]
		index++
		start2++
	}

	// 把临时数组中的元素放回原数组
	copy(nums[start:end+1], tmp)
}
