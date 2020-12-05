package bubblesort

import "testing"

func TestBubbleSort(t *testing.T) {
	testTables := []struct {
		input  []int
		output []int
	}{
		{[]int{2, 3, 1, 5}, []int{1, 2, 3, 5}},
		{[]int{6, 7, 8, 1}, []int{1, 6, 7, 8}},
		{[]int{213, 21421, 22, 111}, []int{22, 111, 213, 21421}},
	}

	for _, testTable := range testTables {
		bubbleSort(testTable.input)
		for index := range testTable.input {
			if testTable.input[index] != testTable.output[index] {
				t.Error("Something is wrong. 😟")
			}
		}
	}
}

func BenchmarkBubbleSort(b *testing.B) {
	for i := 0; i < b.N; i++ {
		bubbleSort([]int{213, 21421, 22, 111, 213, 21421, 22, 111, 12312, 7564, 1, 23, 41, 647,
			213, 21421, 22, 111, 213, 21421, 22, 111, 12312, 7564, 1, 23, 41, 647,
			213, 21421, 22, 111, 213, 21421, 22, 111, 12312, 7564, 1, 23, 41, 647,
			213, 21421, 22, 111, 213, 21421, 22, 111, 12312, 7564, 1, 23, 41, 647,
			213, 21421, 22, 111, 213, 21421, 22, 111, 12312, 7564, 1, 23, 41, 647})
	}
}
