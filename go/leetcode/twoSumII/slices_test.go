package twoSumII

import "testing"

type binarySearchTestCase struct {
	arr      []int
	num      int
	expected int
}

var binarySearchTestCases = []binarySearchTestCase{
	{[]int{}, 0, -1},
	{[]int{}, 1, -1},
	{[]int{-1}, 1, -1},
	{[]int{-1}, -1, 0},
	{[]int{-1, 1}, -1, 0},
	{[]int{-1, 1}, 1, 1},
	{[]int{-1, 0, 1}, -1, 0},
	{[]int{-1, 0, 1}, 0, 1},
	{[]int{-1, 0, 1}, 1, 2},
	{[]int{1, 10, 99, 100}, 2, -1},
	{[]int{1, 10, 99, 100}, 1, 0},
	{[]int{1, 10, 99, 100}, 10, 1},
	{[]int{1, 10, 99, 100}, 99, 2},
	{[]int{1, 10, 99, 100}, 100, 3},
	{[]int{1, 10, 50, 99, 100}, 0, -1},
	{[]int{1, 10, 50, 99, 100}, 1, 0},
	{[]int{1, 10, 50, 99, 100}, 10, 1},
	{[]int{1, 10, 50, 99, 100}, 50, 2},
	{[]int{1, 10, 50, 99, 100}, 99, 3},
	{[]int{1, 10, 50, 99, 100}, 100, 4},
}

func testBinarySearch(t *testing.T, binarySearch func([]int, int) int) {
	for _, tc := range binarySearchTestCases {
		if got := binarySearch(tc.arr, tc.num); got != tc.expected {
			t.Errorf("binarySearch(%v, %d) returned %d, expected %d", tc.arr, tc.num, got, tc.expected)
		}
	}
}
func TestBinarySearch(t *testing.T) {
	testBinarySearch(t, binarySearch)
}

func TestBinarySearchIterative(t *testing.T) {
	testBinarySearch(t, binarySearchRecursive)
}

func TestBinarySearchRecursive(t *testing.T) {
	testBinarySearch(t, binarySearchRecursive)
}

var binarySearchArr []int = []int{1, 1, 11, 22, 33, 44, 55, 66, 77, 88, 99, 100, 101}

func benchmarkBinarySearch(b *testing.B, binarySearch func([]int, int) int, arr []int, num int) {
	for i := 0; i < b.N; i++ {
		_ = binarySearch(arr, num)
	}
}

func BenchmarkBinarySearchLeft(b *testing.B) {
	benchmarkBinarySearch(b, binarySearch, binarySearchArr, 1)
}

func BenchmarkBinarySearchMiddle(b *testing.B) {
	benchmarkBinarySearch(b, binarySearch, binarySearchArr, 55)
}

func BenchmarkBinarySearchRight(b *testing.B) {
	benchmarkBinarySearch(b, binarySearch, binarySearchArr, 101)
}

func BenchmarkBinarySearchRecursiveLeft(b *testing.B) {
	benchmarkBinarySearch(b, binarySearchRecursive, binarySearchArr, 1)
}

func BenchmarkBinarySearchRecursiveMiddle(b *testing.B) {
	benchmarkBinarySearch(b, binarySearchRecursive, binarySearchArr, 55)
}

func BenchmarkBinarySearchRecursiveRight(b *testing.B) {
	benchmarkBinarySearch(b, binarySearchRecursive, binarySearchArr, 101)
}

func BenchmarkBinarySearchIterativeLeft(b *testing.B) {
	benchmarkBinarySearch(b, binarySearchIterative, binarySearchArr, 1)
}

func BenchmarkBinarySearchIterativeMiddle(b *testing.B) {
	benchmarkBinarySearch(b, binarySearchIterative, binarySearchArr, 55)
}

func BenchmarkBinarySearchIterativeRight(b *testing.B) {
	benchmarkBinarySearch(b, binarySearchIterative, binarySearchArr, 101)
}
