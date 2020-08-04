#!/usr/bin/env python3


def add(arr, num=1):
    res, carry = [], 1
    for n in reversed(arr):
        n += carry
        if n == 10:
            carry, n = 1, 0
        else:
            carry = 0
        res.append(n)
    if carry:
        res.append(carry)
    return list(reversed(res))


def test():
    assert add1([]) == [1]
    assert add1([0]) == [1]
    assert add1([0, 0, 0]) == [0, 0, 1]
    assert add1([1, 1, 1]) == [1, 1, 2]
    assert add1([9, 9, 1]) == [9, 9, 2]
    assert add1([9, 9, 9]) == [1, 0, 0, 0]
    print('Successful!')


if __name__ == '__main__':
    test()
