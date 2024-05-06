package interpolation_search

func InterpolationSearch(nums []int, target int) int {
	left, right := 0, len(nums)-1

	if target < nums[left] || target > nums[right] {
		return -1
	}

	for left <= right {
		if left == right {
			if nums[left] == target {
				return left
			} else {
				return -1
			}
		}

		mid := left + (target-nums[left])*(right-left)/(nums[right]-nums[left])
		if nums[mid] == target {
			return mid
		} else if nums[mid] < target {
			left = mid + 1
		} else {
			right = mid - 1
		}
	}

	return -1
}
