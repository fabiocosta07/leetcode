# Definition for a binary tree node.
# class TreeNode:
#     def __init__(self, val=0, left=None, right=None):
#         self.val = val
#         self.left = left
#         self.right = right
class Solution:
    def isValidBST(self, root: Optional[TreeNode]) -> bool:
        self.seq = []
        self.traversal(root)        
        print(self.seq)
    
        for i in range(len(self.seq)):
            if i > 0:
                if not self.seq[i - 1] < self.seq[i]:
                    return False
                    
        return True
    def traversal(self, node: Optional[TreeNode]):
        if node.left:
            self.traversal(node.left)

        self.seq.append(node.val)

        if node.right:
            self.traversal(node.right)
        
        