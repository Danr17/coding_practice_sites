import unittest
from last_stone_weight import Solution

class TestSingleNumber(unittest.TestCase):
    def test_singleNumber(self):
        """
        Given nums = [2,7,4,1,8,1]
        return 1.
        """
        nums = [2,7,4,1,8,1]
        expected = 1

        sol = Solution()
        result = sol.lastStoneWeight(nums)
        self.assertEqual(int(result), expected)

if __name__ == '__main__':
    unittest.main()