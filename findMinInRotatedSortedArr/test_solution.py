import unittest

from solution import Solution


class TestSol(unittest.TestCase):

    def test_sol1(self):
        sol = Solution()
        res = sol.findMin([3,4,5,1,2])
        self.assertEqual(res, 1, "Should be 1")
    def test_sol2(self):
        sol = Solution()
        res = sol.findMin([4,5,6,7,0,1,2])
        self.assertEqual(res, 0, "Should be 0")
    def test_sol3(self):
        sol = Solution()
        res = sol.findMin([11,13,15,17])
        self.assertEqual(res, 11, "Should be 11")
    def test_sol4(self):
        sol = Solution()
        res = sol.findMin([2,1])
        self.assertEqual(res, 1, " [2, 1] Should be 1")
    def test_sol5(self):
        sol = Solution()
        res = sol.findMin([1])
        self.assertEqual(res, 1, " [1] Should be 1")
    def test_sol6(self):
        sol = Solution()
        res = sol.findMin([2,3,4,5,1])
        self.assertEqual(res, 1, " [2,3,4,5,1] Should be 1")
    def test_sol7(self):
        sol = Solution()
        res = sol.findMin([3,1,2])
        self.assertEqual(res, 1, " [3,1,2] Should be 1")



if __name__ == '__main__':
    unittest.main()