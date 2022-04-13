# Definition for singly-linked list.
from typing import Optional


class ListNode:
     def __init__(self, val=0, next=None):
         self.val = val
         self.next = next
class Solution:
    def addTwoNumbers(self, l1: Optional[ListNode], l2: Optional[ListNode]) -> Optional[ListNode]:
        cur1 = l1
        cur2 = l2
        res  = None
        resHead = None
        carie = False
        sum = 0
        while (cur1 != None or cur2 != None or carie):
          if cur1 != None:
            sum += cur1.val
          if cur2 != None:
            sum += cur2.val

          if carie:
            sum += 1
            carie = False

          if sum > 9:
            sum -= 10            
            carie = True          

          if res == None:
            res = ListNode(sum)
            resHead = res
          else:
            res.next = ListNode(sum)
            res = res.next
          sum = 0
          if cur1 != None:
            cur1 = cur1.next
          if cur2 != None:
            cur2 = cur2.next

        return resHead
          
          