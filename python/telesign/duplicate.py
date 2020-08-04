"""
Let's say you have a list of N+1 integers between 1 and N. You know there's at
least one duplicate, but there might be more. For example, if N=3, your list
might be 3, 1, 1, 3 or it might be 1, 3, 2, 2. Print out a number that appears
in the list more than once. (That is, in the first example, you can print '1'
or '3' -- you don't have to print both.)
"""

import collections as col
import random


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
    # complexity: O(nlgn)
    # memory: O(1) additional memory
    n = len(integers) - 1
    lo, hi = 1, n
    while lo < hi:  # lgn iterations * sum of n values
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
    print('test success')
    assert dupe([3, 1, 1, 3]) in [1, 3]
    print('test success')
    assert dupe([1, 1, 2, 2]) in [1, 2]
    print('test success')
    assert dupe([1, 1, 2, 2, 3]) in [1, 2]
    print('test success')
    assert dupe([1, 1, 2, 4, 4]) in [1, 4]
    print('test success')
    assert dupe([1, 2, 3, 4, 4]) == 4
    print('test success')
    assert dupe([1, 1]) == 1
    print('test success')
    assert dupe([1, 1, 2]) == 1
    print('test success')
    assert dupe([1, 2, 2]) == 2
    print('test success')
    assert dupe([1, 1, 1]) == 1
    print('test success')
    assert dupe([2, 2, 2]) == 2
    print('test success')
    test = list(range(1, 100))
    for x in test:
        mytest = list(test)
        random.shuffle(mytest)
        for _ in range(x - 1):
            mytest.pop()
        mytest = mytest + [x] * x
        random.shuffle(mytest)
        assert dupe(mytest) == x
        print('test success %d' % x)
