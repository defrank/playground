package twoSumII

import "testing"

type twoSumTestCase struct {
	numbers  []int
	target   int
	expected []int
}

var twoSumTestCases = []twoSumTestCase{
	{[]int{}, 0, nil},
	{[]int{}, 1, nil},
	{[]int{0}, 0, nil},
	{[]int{1, 1}, 1, nil},
	{[]int{2, 7, 11, 15}, 1, nil},
	{[]int{2, 7, 11, 15}, 9, []int{1, 2}},
	{[]int{1, 1, 2, 4, 5, 6, 9}, 9, []int{4, 5}},
	{[]int{1, 2, 3, 4, 5}, 9, []int{4, 5}},
	{[]int{1, 2, 3, 4, 6}, 6, []int{2, 4}},
}

func testTwoSum(t *testing.T, twoSum func([]int, int) []int) {
	failure := func(t *testing.T, tc twoSumTestCase, got []int) {
		t.Errorf("twoSum(%v, %d) returned %v, expected %v", tc.numbers, tc.target, got, tc.expected)
	}

	for _, tc := range twoSumTestCases {
		if got := twoSum(tc.numbers, tc.target); tc.expected == nil && got != nil {
			failure(t, tc, got)
		} else if len(got) != len(tc.expected) {
			failure(t, tc, got)
		} else {
			for i, v := range got {
				if v != tc.expected[i] {
					failure(t, tc, got)
				}
			}
		}
	}
}

func TestTwoSum(t *testing.T) {
	testTwoSum(t, twoSum)
}

func TestTwoSumBruteForce(t *testing.T) {
	testTwoSum(t, twoSumBruteForce)
}

func TestTwoSumBinarySearch(t *testing.T) {
	testTwoSum(t, twoSumBinarySearch)
}

func TestTwoSumBounded(t *testing.T) {
	testTwoSum(t, twoSumBounded)
}

var twoSumBenchNumbers []int = []int{1, 1, 2, 3, 5, 5}

func benchmarkTwoSum(b *testing.B, twoSum func([]int, int) []int, numbers []int, target int) {
	for i := 0; i < b.N; i++ {
		_ = twoSum(numbers, target)
	}
}

func BenchmarkTwoSumLeft(b *testing.B) {
	benchmarkTwoSum(b, twoSum, twoSumBenchNumbers, 2)
}

func BenchmarkTwoSumMiddle(b *testing.B) {
	benchmarkTwoSum(b, twoSum, twoSumBenchNumbers, 5)
}

func BenchmarkTwoSumRight(b *testing.B) {
	benchmarkTwoSum(b, twoSum, twoSumBenchNumbers, 10)
}

func BenchmarkTwoSumBinarySearchLeft(b *testing.B) {
	benchmarkTwoSum(b, twoSumBinarySearch, twoSumBenchNumbers, 2)
}

func BenchmarkTwoSumBinarySearchMiddle(b *testing.B) {
	benchmarkTwoSum(b, twoSumBinarySearch, twoSumBenchNumbers, 5)
}

func BenchmarkTwoSumBinarySearchRight(b *testing.B) {
	benchmarkTwoSum(b, twoSumBinarySearch, twoSumBenchNumbers, 10)
}

func BenchmarkTwoSumBoundedLeft(b *testing.B) {
	benchmarkTwoSum(b, twoSumBounded, twoSumBenchNumbers, 2)
}

func BenchmarkTwoSumBoundedMiddle(b *testing.B) {
	benchmarkTwoSum(b, twoSumBounded, twoSumBenchNumbers, 5)
}

func BenchmarkTwoSumBoundedRight(b *testing.B) {
	benchmarkTwoSum(b, twoSumBounded, twoSumBenchNumbers, 10)
}
