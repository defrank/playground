"""
Given the function signature:

    def combine(abs_path, paths)

Input:

    "/a/b/c/d", "../e/f/../../g/./i"

Output:

    "/a/b/c/g/i"
"""
import contextlib as cl
import itertools as it


SEPARATOR = '/'
PARENT = '..'
CURRENT = '.'
EMPTY = ''


def combine(abs_path, paths):
    stack = []
    filenames = it.chain.from_iterable(
        path.strip(SEPARATOR).split(SEPARATOR)
        for path in (abs_path, paths)
    )
    for filename in filenames:
        if filename == CURRENT or filename == EMPTY:
            continue
        elif filename == PARENT:
            with cl.suppress(IndexError):
                stack.pop()
        else:
            stack.append(filename)

    return SEPARATOR + SEPARATOR.join(stack)


if __name__ == '__main__':
    assert combine('/a/b/c/d', '../e/f/../../g/./i') == '/a/b/c/g/i'
    assert combine('/', '..') == '/'
    assert combine('/', '../..') == '/'
    assert combine('/a/b/', 'c/d') == '/a/b/c/d'
    assert combine('/a/b/', '.') == '/a/b'
    assert combine('/a/b/', './.') == '/a/b'
    assert combine('//', '.') == '/'
    print('success!')
