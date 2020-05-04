from typing import List

class Solution:
    def productExceptSelf(self, nums: List[int]) -> List[int]:
        result = [1 for _ in nums]
        for start, end, dir in [(1, len(nums), 1), (len(nums)-2, -1, -1)]:
            cur = 1
            for i in range(start, end, dir):
                cur *= nums[i-dir]
                result[i] *= cur
        return result

