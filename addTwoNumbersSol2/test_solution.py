from operator import add
from typing import List
import unittest

from solution import ListNode, Solution


class TestSol(unittest.TestCase):

    def test_1(self):
        l1 = helperConvertArrToLinkedList([2,4,3])
        l2 = helperConvertArrToLinkedList([5,6,4])        
        sol = Solution()
        res = sol.addTwoNumbers(l1, l2)
        testRes = helperConvertLinkedListToArr(res)
        self.assertEqual(testRes, [7,0,8], "Should be 807")

def helperConvertArrToLinkedList(l: List) -> ListNode:
    ll = None
    llhead = None
    for n in l:
        if ll == None:
          ll = ListNode(n)
          llhead = ll
        else:
          ll.next = ListNode(n)
          ll = ll.next
    return llhead

def helperConvertLinkedListToArr(l: ListNode) -> ListNode:
    lres = []
    cur = l
    while(cur != None):
      lres.append(cur.val) 
      cur = cur.next    
    return lres


if __name__ == '__main__':
    unittest.main()