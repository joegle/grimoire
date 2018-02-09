#!/usr/bin/env python2

# return the sorted array
a = [2, 3, 1]
print sorted(a)

# in place sort returning None
a.sort()

# sort by a function key
a.sort(key=str.lower, reverse=True)
