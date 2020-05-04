import unittest
from backspace_string_compare import Solution

class TestSingleNumber(unittest.TestCase):
    def test_first(self):

        S = "ab#c"
        T = "ad#c"
        expected = True

        sol = Solution()
        result = sol.backspaceCompare(S, T)
        self.assertEqual(result, expected)
    
    def test_second(self):

        S = "a##c"
        T = "#a#c"
        expected = True

        sol = Solution()
        result = sol.backspaceCompare(S, T)
        self.assertEqual(result, expected)
    
    def test_Third(self):

        S = "a#c"
        T = "b"
        expected = False

        sol = Solution()
        result = sol.backspaceCompare(S, T)
        self.assertEqual(result, expected)


if __name__ == '__main__':
    unittest.main()