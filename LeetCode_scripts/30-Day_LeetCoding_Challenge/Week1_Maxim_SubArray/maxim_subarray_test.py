import unittest
from maxim_subarray import Solution

class TestSingleNumber(unittest.TestCase):
    def test_singleNumber(self):
        """
        Given nums = [-2,1,-3,4,-1,2,1,-5,4]
        return 6 , which is [4,-1,2,1]
        """
        nums = [-2,1,-3,4,-1,2,1,-5,4]
        expected = 6

        sol = Solution()
        result = sol.maxSubArray(nums)
        self.assertEqual(result, expected)

if __name__ == '__main__':
    unittest.main()