import unittest
from group_anagrams import Solution

class SolutionTest(unittest.TestCase):
    def test_anagrams(self):
        """
        given: ["eat", "tea", "tan", "ate", "nat", "bat"],
        return: 
        [
            ["ate","eat","tea"],
            ["nat","tan"],
            ["bat"]
        ]
        """
        theList = ["eat", "tea", "tan", "ate", "nat", "bat"]
        expected = [
            ["eat","tea","ate"],
            ["tan", "nat"],
            ["bat"]
        ]

        sol = Solution()
        result = sol.groupAnagrams(theList)
        self.assertEqual(result, expected)

if __name__ == '__main__':
    unittest.main()