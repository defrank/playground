"""
Provided
* a list of integers
* target value

Goal:
* summing of subset of list equals the target


Examples:

    * [1, 2, 3, 4, 5], 6, min={1}=1, max={1,2,3,4,5}=15
    * [1, 2, 3, 4, 5], 11, min={1}=0, max={1,2,3,4,5}=15 => {2,4,5}, {1,2,3,5}

"""
from typing import Iterable


def subset_sum(iterable: Iterable, target: int) -> bool:
    numbers = list(iterable)

    if sum(numbers) == target:
        return True
    elif not numbers:
        return False
    return any(
        subset_sum(numbers[:idx] + numbers[idx+1:], target)
        for idx in range(len(numbers))
    )


def subset_sum_positive(iterable: Iterable, target: int) -> bool:
    numbers = list(iterable)

    if target == 0:
        return True
    elif target < 0:
        return False
    elif not numbers:
        return False
    return (
        subset_sum_positive(numbers[1:], target - numbers[0])
        or subset_sum_positive(numbers[1:], target)
    )


def subset_sum2(iterable: Iterable, target: int) -> bool:
    """
    O(n!)
    """
    numbers = list(iterable)

    if target == 0:
        return True
    elif not numbers:
        return False
    return (
        subset_sum2(numbers[1:], target - numbers[0])
        or subset_sum2(numbers[1:], target)
    )
