import unittest

from two_sum import Solution

class TestTwoSum(unittest.TestCase):
    def test_sum(self):
        """
        Given nums = [2, 7, 11, 15], target = 9,
        Because nums[0] + nums[1] = 2 + 7 = 9,
        return [0, 1].
        """

        nums = [2, 7, 11, 15]
        target = 9
        expected = [0, 1]

        sol = Solution()
        result = sol.twoSum(nums, target)
        self.assertListEqual(result, expected)

if __name__ == '__main__':
    unittest.main()