package twoSumII

/**
 * Return the 0-based index of num found in arr if found; otherwise -1
 */
func binarySearch(arr []int, num int) int {
	return binarySearchIterative(arr, num)
}

func binarySearchIterative(arr []int, num int) int {
	// n = len(arr)
	// complexity: O(lg(n))
	// memory: O(n)
	left, right := 0, len(arr)-1
	for left <= right {
		idx := (left + right) / 2
		if val := arr[idx]; val > num {
			right = idx - 1
		} else if val < num {
			left = idx + 1
		} else {
			return idx
		}
	}
	return -1
}

func binarySearchRecursive(arr []int, num int) int {
	// n = len(arr)
	// complexity: O(lg(n))
	// memory: O(n)
	if len(arr) == 0 {
		return -1
	}

	idx := len(arr) / 2
	if val := arr[idx]; val > num {
		idx = binarySearch(arr[:idx], num)
	} else if val < num {
		var relativeIdx int
		relativeIdx = binarySearch(arr[idx+1:], num)
		if relativeIdx == -1 {
			idx = -1
		} else {
			idx += relativeIdx + 1
		}
	}
	return idx
}
