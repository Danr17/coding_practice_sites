from typing import List

class Solution:
    def moveZeroes(self, nums: List[int]) -> None:
        """
        Do not return anything, modify nums in-place instead.
        """
        for i, n in enumerate(nums):
            if n == 0:
                nums.append(nums[i])
                nums.remove(nums[i])
                i -= 1

