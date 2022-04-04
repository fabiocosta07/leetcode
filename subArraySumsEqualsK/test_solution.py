import unittest

from solution import Solution


class TestSol(unittest.TestCase):

    def test_sol1(self):
        sol = Solution()
        res = sol.subarraySum([1,1,1], 2)
        self.assertEqual(res, 2, "Should be 2")
    def test_sol2(self):
        sol = Solution()
        res = sol.subarraySum([1,2,3], 3)
        self.assertEqual(res, 2, "Should be 2")
    def test_sol3(self):
        sol = Solution()
        res = sol.subarraySum([1], 0)
        self.assertEqual(res, 0, "Should be 0")
    def test_sol4(self):
        sol = Solution()
        res = sol.subarraySum([-1,-1,1], 0)
        self.assertEqual(res, 1, "Should be 1")
    def test_sol5(self):
        sol = Solution()
        res = sol.subarraySum([1,-1,0], 0)
        self.assertEqual(res, 3, "Should be 3")

if __name__ == '__main__':
    unittest.main()