import unittest
from counting_elements import Solution

class SolutionTest(unittest.TestCase):
    def test_growing_numbers(self):
        """
        given: [1,2,3]
        return: 2
        """
        theList = [1,2,3]
        expected = 2

        sol = Solution()
        result = sol.countElements(theList)
        self.assertEqual(result, expected)
    
    def test_zero_value(self):
        """
        given: [1,1,3,3,5,5,7,7]
        return: 0
        """
        theList = [1,1,3,3,5,5,7,7]
        expected = 0

        sol = Solution()
        result = sol.countElements(theList)
        self.assertEqual(result, expected)
    
    def test_unorder_numbers(self):
        """
        given: [1,3,2,3,5,0]
        return: 3
        """
        theList = [1,3,2,3,5,0]
        expected = 3

        sol = Solution()
        result = sol.countElements(theList)
        self.assertEqual(result, expected)
    
    def test_double_values(self):
        """
        given: [1,1,2,2]
        return: 2
        """
        theList = [1,1,2,2]
        expected = 2

        sol = Solution()
        result = sol.countElements(theList)
        self.assertEqual(result, expected)

if __name__ == '__main__':
    unittest.main()