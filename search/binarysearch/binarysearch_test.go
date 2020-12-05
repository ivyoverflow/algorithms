package binarysearch

import (
	"testing"
)

func TestBinarySearch(t *testing.T) {
	testTables := []struct {
		input  []int
		target int
		result int
	}{
		{[]int{1, 2, 3, 4, 5}, 1, 0},
		{[]int{1, 6, 7, 8}, 8, 3},
		{[]int{22, 111, 213, 21421}, 111, 1},
	}

	for _, testTable := range testTables {
		result := binarySearch(testTable.input, testTable.target)
		if result != testTable.result {
			t.Error("Something is wrong. 😟")
		}
	}
}

func BenchmarkBinarySearch(b *testing.B) {
	for i := 0; i < b.N; i++ {
		binarySearch([]int{213, 21421, 22, 111, 213, 21421, 22, 111, 12312, 7564, 1, 23, 41, 647}, 1)
	}
}
