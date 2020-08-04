package twoSumII

const IndexOffset = 1

/**
 * Return 1-based indices of values that sum to the target or nil.
 * Assumes input numbers are sorted.
 *
 * https://leetcode.com/problems/two-sum-ii-input-array-is-sorted/
 */
func twoSum(numbers []int, target int) []int {
	return twoSumBounded(numbers, target)
}

func twoSumBounded(numbers []int, target int) []int {
	// n = len(numbers)
	// complexity: O(n)
	// memory: O(n)
	leftIdx, rightIdx := 0, len(numbers)-1
	for leftIdx < rightIdx {

		if sum := numbers[leftIdx] + numbers[rightIdx]; sum < target {
			leftIdx++
		} else if sum > target {
			rightIdx--
		} else {
			return []int{leftIdx + IndexOffset, rightIdx + IndexOffset}
		}
	}
	return nil
}

func twoSumBinarySearch(numbers []int, target int) []int {
	// n = len(numbers)
	// complexity: O(n*binarySearch) -> O(n*lg(n))
	// memory: O(n)
	for leftIdx, leftVal := range numbers {
		rightIdx := leftIdx + 1
		relativeIdx := binarySearch(numbers[rightIdx:], target-leftVal)
		if relativeIdx != -1 {
			return []int{leftIdx + IndexOffset, rightIdx + relativeIdx + IndexOffset}
		}
	}

	return nil
}

func twoSumBruteForce(numbers []int, target int) []int {
	// n = len(numbers)
	// complexity: O(n*(n+1)/2) -> O(n^2)
	// memory: O(n)
	for leftIdx, leftVal := range numbers {
		toFind := target - leftVal
		for rightIdx, rightVal := range numbers[leftIdx+1:] {
			if rightVal == toFind {
				return []int{
					leftIdx + IndexOffset,
					leftIdx + rightIdx + 1 + IndexOffset,
				}
			}
		}
	}
	return nil
}
