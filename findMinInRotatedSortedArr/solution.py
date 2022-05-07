from typing import List


class Solution:
    def findMin(self, nums: List[int]) -> int:        
        if len(nums) == 2:
          return min(nums[0], nums[1])
        if len(nums) == 3:
          return min(nums[0], nums[1], nums[2])
        mi = self.getMiddleIndex(len(nums))        
        while nums[0] > nums[len(nums) - 1]:
          if nums[mi] >= nums[0]:
            nums = nums[mi + 1:]
          else:
            nums = nums[:mi]
          mi = self.getMiddleIndex(len(nums))

        return nums[0]

    def getMiddleIndex(self, size: int) -> int:
        if size % 2 == 0:
          return int (size / 2)
        else:
          return int((size - 1) / 2)