package selectionsort

func selectionSort(nums []int) {
	for i := 0; i < len(nums); i++ {
		minIndex := i
		for j := i + 1; j < len(nums); j++ {
			if nums[j] < nums[minIndex] {
				minIndex = j
			}
		}

		nums[i], nums[minIndex] = nums[minIndex], nums[i]
	}
}
