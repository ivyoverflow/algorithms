package insertionsort

func insertionSort(nums []int) {
	for i := 1; i < len(nums); i++ {
		for j := 0; j < i; j++ {
			if nums[i] < nums[j] {
				nums[i], nums[j] = nums[j], nums[i]
			}
		}
	}
}
