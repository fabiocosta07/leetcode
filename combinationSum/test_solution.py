from typing import List
import unittest

from solution import Solution


class TestSol(unittest.TestCase):

    def test_sol1(self):
        sol = Solution()
        res = sol.combinationSum([2,3,6,7], 7)        
        self.assertTrue(isEqual(res, [[2,2,3],[7]]), "Should be [[2,2,3],[7]]")

def isEqual(l1: List[List[int]], l2: List[List[int]]) -> bool:
    for c1 in l1:
        if not hasCombinationSum(c1, l2):
            return False
    return True
def hasCombinationSum(c1: List[int], l: List[List[int]]) -> bool:
    for c in l:
        if sum(c) == sum(c1):
            return True
    return False

if __name__ == '__main__':
    unittest.main()