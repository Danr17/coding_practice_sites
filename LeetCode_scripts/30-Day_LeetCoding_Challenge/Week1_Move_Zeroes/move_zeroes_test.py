import unittest
from move_zeroes import Solution

class TestSingleNumber(unittest.TestCase):
    def test_singleNumber(self):
        """
        Given nums = [-2,1,-3,4,-1,2,1,-5,4]
        return 6 , which is [4,-1,2,1]
        """
        nums = [0,1,0,3,12]
        expected = [1,3,12,0,0]

        sol = Solution()
        sol.moveZeroes(nums)
        self.assertEqual(nums, expected)

if __name__ == '__main__':
    unittest.main()