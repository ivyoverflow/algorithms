package linearsearch

import "testing"

func TestLinearSearch(t *testing.T) {
	testTables := []struct {
		input  []int
		target int
		result int
	}{
		{[]int{2, 3, 1, 5}, 1, 2},
		{[]int{6, 7, 8, 1}, 8, 2},
		{[]int{213, 21421, 22, 111}, 111, 3},
	}

	for _, testTable := range testTables {
		result := linearSearch(testTable.input, testTable.target)
		if result != testTable.result {
			t.Error("Something is wrong. 😟")
		}
	}
}

func BenchmarkLinearSearch(b *testing.B) {
	for i := 0; i < b.N; i++ {
		linearSearch([]int{213, 21421, 22, 111, 213, 21421, 22, 111, 12312, 7564, 1, 23, 41, 647}, 1)
	}
}
