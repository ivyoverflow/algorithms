package quicksort

import (
	"math/rand"
)

func quickSort(nums []int) []int {
	if len(nums) <= 1 {
		return nums
	}

	median := nums[rand.Intn(len(nums))]
	var low []int
	var middle []int
	var high []int

	for _, num := range nums {
		switch {
		case num < median:
			low = append(low, num)
		case num == median:
			middle = append(middle, num)
		default:
			high = append(high, num)
		}
	}

	low = quickSort(low)
	high = quickSort(high)

	low = append(low, middle...)
	low = append(low, high...)

	return low
}
