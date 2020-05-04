from typing import List

class Solution:
    def singleNumber(self, nums: List[int]) -> int:
        s = {}
        for n in nums:
            if n in s:
                del s[n]
                continue
            s[n] = True
        return " ".join(str(key) for key in s.keys())