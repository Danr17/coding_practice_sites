from typing import List

class Solution:
    def twoSum(self, nums: List[int], target: int) -> List[int]:
        position = {}

        for i, num in enumerate(nums):
            if (target - nums[i]) in position.keys():
                result = [position[target - nums[i]], i]
                return result

            position[num] = i
        
        return []
