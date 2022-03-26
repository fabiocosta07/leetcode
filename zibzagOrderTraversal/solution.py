# Definition for a binary tree node.
# class TreeNode:
#     def __init__(self, val=0, left=None, right=None):
#         self.val = val
#         self.left = left
#         self.right = right
import collections


class Solution:
    def zigzagLevelOrder(self, root: Optional[TreeNode]) -> List[List[int]]:
        res = []

        q = collections.deque()
        q.append(root)

        toogle = True
        while q:
            qLen = len(q)
            level = []
            for i in range(qLen):
                if toogle: 
                    node = q.popleft()
                    if node: 
                        level.append(node.val)
                        q.append(node.left)
                        q.append(node.right)
                else:
                    node = q.pop()
                    if node: 
                        level.append(node.val)
                        q.appendleft(node.right)
                        q.appendleft(node.left)
            toogle = not toogle
            if level:
                res.append(level)

        return res