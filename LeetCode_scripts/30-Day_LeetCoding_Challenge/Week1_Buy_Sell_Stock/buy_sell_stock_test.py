import unittest
from buy_sell_stock import Solution

class SolutionTest(unittest.TestCase):
    def test_normal(self):
        """
        given: [7,1,5,3,6,4]
        return: 7
        """
        num = [7,1,5,3,6,4]
        expected = 7

        sol = Solution()
        self.assertEqual(sol.maxProfit(num), expected)
    
    def test_growing_numbers(self):
        """
        given: [1,2,3,4,5]
        return: 4
        """
        num = [1,2,3,4,5]
        expected = 4

        sol = Solution()
        self.assertEqual(sol.maxProfit(num), expected)
    
    def test_shrinking_numbers(self):
        """
        given: [7,6,4,3,1]
        return: 0
        """
        num = [7,6,4,3,1]
        expected = 0

        sol = Solution()
        self.assertEqual(sol.maxProfit(num), expected)

if __name__ == '__main__':
    unittest.main()