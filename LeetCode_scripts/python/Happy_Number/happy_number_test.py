import unittest
from happy_number import Solution

class SolutionTest(unittest.TestCase):
    def test_isHappy(self):
        """
        given: 19
        return: True
        """
        num = 19
        expected = True

        sol = Solution()
        result = sol.isHappy(num)
        self.assertIs(result, expected)

if __name__ == '__main__':
    unittest.main()