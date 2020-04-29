# Definition for a binary tree node.
class TreeNode:
     def __init__(self, x):
         self.val = x
         self.left = None
         self.right = None

class Solution:
    def diameterOfBinaryTree(self, root: TreeNode) -> int:
        self.res = 0
        def dfs(node):
            if node == None:
                return 0
            l = dfs(node.left)
            r = dfs(node.right)
            self.res = max(self.res, l + r)
            self.res = max(self.res, max(l, r))
            return max(l, r) + 1
        
        dfs(root)
        return self.res