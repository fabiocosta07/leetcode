import unittest

from solution import Solution


class TestSol(unittest.TestCase):

    def test_sol1(self):
        sol = Solution()
        res = sol.combinationSum([2,3,6,7], 7)
        self.assertEqual(res, 1, "Should be ")

if __name__ == '__main__':
    unittest.main()