import unittest
from single_number import Solution

class TestSingleNumber(unittest.TestCase):
    def test_singleNumber(self):
        """
        Given nums = [4,1,2,1,2]
        return 4.
        """
        nums = [4,1,2,1,2]
        expected = 4

        sol = Solution()
        result = sol.singleNumber(nums)
        self.assertEqual(int(result), expected)

if __name__ == '__main__':
    unittest.main()