import unittest
from contiguous_array import Solution

class TestSingleNumber(unittest.TestCase):
    def test_singleNumber(self):
        """
        Input:  [0,1,0]
        Output: 2
        """
        nums = [0,1,0]
        expected = 2

        sol = Solution()
        result = sol.findMaxLength(nums)
        self.assertEqual(result, expected)

if __name__ == '__main__':
    unittest.main()