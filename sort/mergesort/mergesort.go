package mergesort

func merge(left, right []int) (out []int) {
	if len(left) > 0 || len(right) > 0 {
		switch {
		case len(left) == 0:
			out = append(out, right...)
		case len(right) == 0:
			out = append(out, left...)
		case left[0] <= right[0]:
			out = append(out, left[0])
			left = left[1:]
		default:
			out = append(out, right[0])
			right = right[1:]
		}
	}

	return out
}

func mergeSort(nums []int) []int {
	if len(nums) <= 1 {
		return nums
	}

	middle := len(nums) / 2
	left := mergeSort(nums[:middle])
	right := mergeSort(nums[middle:])

	return merge(left, right)
}
