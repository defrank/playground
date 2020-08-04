package twoSum

const IndexOffset = 0

/**
 * Return 0-based indices of values that sum to the target or nil.
 *
 * https://leetcode.com/problems/two-sum/
 */
func twoSum(nums []int, target int) []int {
	return twoSumBruteForce(nums, target)
}

func twoSumBruteForce(nums []int, target int) []int {
	// n = len(nums)
	// complexity: O(n*(n+1)/2) -> O(n^2)
	// memory: O(n)
	for leftIdx, leftVal := range nums {
		toFind := target - leftVal
		for rightIdx, rightVal := range nums[leftIdx+1:] {
			if rightVal == toFind {
				return []int{leftIdx, leftIdx + rightIdx + 1}
			}
		}
	}
	return nil
}
