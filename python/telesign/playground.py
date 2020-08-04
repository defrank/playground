"""
Let's say you have a list of N+1 integers between 1 and N. You know there's at
least one duplicate, but there might be more. For example, if N=3, your list
might be 3, 1, 1, 3 or it might be 1, 3, 2, 2. Print out a number that appears
in the list more than once. (That is, in the first example, you can print '1'
or '3' -- you don't have to print both.)
"""

import collections as col


def print_dupe(integers):
    # complexity: O(n)
    # memory: O(n)
    counts = col.defaultdict(int)
    for x in integers:
        counts[x] += 1
        if counts[x] > 1:
            print(x)
            break


def dupe(integers):
    n = len(integers) - 1
    lo, hi = 1, n
    while lo < hi:
        pivot = int((hi + lo)/2)
        if sum(1 for x in integers if lo <= x < pivot) > (pivot - lo):
            if hi == pivot:
                hi -= 1
            else:
                hi = pivot
        elif sum(1 for x in integers if pivot < x <= hi) > (hi - pivot):
            if lo == pivot:
                lo += 1
            else:
                lo = pivot
        elif sum(1 for x in integers if x == pivot) > 1:
            return pivot
        else:
            raise AssertionError(f'Unexpected case: {lo}, {pivot}, {hi}')
    return lo


if __name__ == '__main__':
    print('begin tests')
    assert dupe([1, 3, 2, 2]) == 2
    print('test 1 success')
    assert dupe([3, 1, 1, 3]) in [1, 3]
    print('test 2 success')
    assert dupe([1, 1, 2, 2]) in [1, 2]
    print('test 3 success')
    assert dupe([1, 1, 2, 2, 3]) in [1, 2]
    print('test 4 success')
    assert dupe([1, 1, 2, 4, 4]) in [1, 4]
    print('test 5 success')
    assert dupe([1, 2, 3, 4, 4]) == 4
    print('test 6 success')
    assert dupe([1, 1]) == 1
    print('test 7 success')
    assert dupe([1, 1, 2]) == 1
    print('test 8 success')
    assert dupe([1, 2, 2]) == 2
    print('test 9 success')
    assert dupe([1, 1, 1]) == 1
    print('test 10 success')
    assert dupe([2, 2, 2]) == 2
    print('test 11 success')
