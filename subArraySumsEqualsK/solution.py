from typing import List


class Solution:
    def subarraySum(self, nums: List[int], k: int) -> int:
        # 0 prefix added for case of sum = k in the firt loop run                
        prefixDict = { 0 : 1}
        sumCounter = 0
        sum = 0
        for i in range(len(nums)):
            sum += nums[i]
            diff = sum - k
            if diff in prefixDict.keys():
                sumCounter+=prefixDict.get(diff)            
            if sum in prefixDict.keys():
                prefixCounter = prefixDict.get(sum)
            else:
                prefixCounter = 0
            prefixCounter+=1
            prefixDict[sum] = prefixCounter
                                    
        return sumCounter
    