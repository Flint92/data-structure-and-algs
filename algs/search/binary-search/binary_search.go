package binary_search

func BinarySearch(nums []int, item int) int {
	left, right := 0, len(nums)-1
	for left <= right {
		mid := left + (right-left)>>1
		if item == nums[mid] {
			return mid
		} else if item < nums[mid] {
			right = mid - 1
		} else {
			left = mid + 1
		}
	}
	return -1
}
