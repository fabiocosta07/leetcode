from typing import List


class Solution:
    def findMin(self, nums: List[int]) -> int:        
        mi = self.getMiddleIndex(len(nums))        
        while len(nums) > 2:
          if self.isSorted(nums):
            return nums[0]
          if nums[mi] >= nums[0]:
            nums = nums[mi + 1:]
          else:
            nums = nums[:mi + 1]
          mi = self.getMiddleIndex(len(nums))

        if len(nums) == 2:
          return min(nums[0], nums[1])

        return nums[0]

    def getMiddleIndex(self, size: int) -> int:
        if size % 2 == 0:
          return int (size / 2)
        else:
          return int((size - 1) / 2)
    def isSorted(self, nums: List[int]) -> bool:      
      return (len(nums) > 0) and (nums[0] <= nums[len(nums) - 1])
