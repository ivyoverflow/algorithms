package binarysearch

func binarySearch(nums []int, target int) int {
	left := 0
	right := len(nums) - 1

	for left <= right {
		index := (left + right) / 2
		switch {
		case nums[index] == target:
			return index
		case nums[index] < target:
			left = index + 1
		default:
			right = index - 1
		}
	}

	return -1
}
