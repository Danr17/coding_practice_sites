import unittest
from strings_shifts import Solution

class TestStringsShift(unittest.TestCase):
    def test_shift_simple(self):
        """
        Input:  s = "abc", shift = [[0,1],[1,2]]
        Output: "cab"
        """
        s = "abc"
        shift = [[0,1],[1,2]]
        expected = "cab"

        sol = Solution()
        result = sol.stringShift(s, shift)
        self.assertEqual(result, expected)
    
    def test_shift_complex(self):
        """
        Input:  s = "abcdefg", shift = [[1,1],[1,1],[0,2],[1,3]]
        Output: "efgabcd"
        """
        s = "abcdefg"
        shift = [[1,1],[1,1],[0,2],[1,3]]
        expected = "efgabcd"

        sol = Solution()
        result = sol.stringShift(s, shift)
        self.assertEqual(result, expected)

if __name__ == '__main__':
    unittest.main()