from typing import List

class Solution:
    def countElements(self, arr: List[int]) -> int:
        hashElem = {}
        result = 0
        for n in sorted(arr):
            if n not in hashElem.keys():
                hashElem[n] = 1
            else:
                hashElem[n] += 1

            if n -1 in hashElem.keys():
                result += hashElem[n-1]
                hashElem[n-1] = 0
        return result


