package quicksort

import (
	"testing"
)

func TestQuickSort(t *testing.T) {
	testTables := []struct {
		input  []int
		output []int
	}{
		{[]int{2, 3, 1, 5}, []int{1, 2, 3, 5}},
		{[]int{6, 7, 8, 1}, []int{1, 6, 7, 8}},
		{[]int{213, 21421, 22, 111}, []int{22, 111, 213, 21421}},
	}

	for _, testTable := range testTables {
		output := quickSort(testTable.input)
		for index := range output {
			if output[index] != testTable.output[index] {
				t.Error("Something is wrong. 😟")
			}
		}
	}
}

func BenchmarkQuickSort(b *testing.B) {
	for i := 0; i < b.N; i++ {
		quickSort([]int{213, 21421, 22, 111, 213, 21421, 22, 111, 12312, 7564, 1, 23, 41, 647})
	}
}
