import unittest
from product_array import Solution

class TestSingleNumber(unittest.TestCase):
    def test_singleNumber(self):
        """
        Input:  [1,2,3,4]
        Output: [24,12,8,6]
        """
        nums = [1,2,3,4]
        expected = [24,12,8,6]

        sol = Solution()
        result = sol.productExceptSelf(nums)
        self.assertEqual(result, expected)

if __name__ == '__main__':
    unittest.main()