from typing import List

class Solution:
    def maxSubArray(self, nums: List[int]) -> int:
        prev = nums[0]
        max_val = prev
        
        for i in range(1, len(nums)):
        
            prev = max(prev + nums[i], nums[i])
            max_val = max(prev, max_val)
                
        return max_val




        