package binary_search

func BinarySearch(nums []int, target int) int {
	left, right := 0, len(nums)-1
	if target < nums[left] || target > nums[right] {
		return -1
	}

	for left <= right {
		mid := left + (right-left)>>1
		if target == nums[mid] {
			return mid
		} else if target < nums[mid] {
			right = mid - 1
		} else {
			left = mid + 1
		}
	}
	return -1
}
