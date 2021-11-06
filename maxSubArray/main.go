package main

import "fmt"

func main() {
	input := []int{1, 2}
	fmt.Printf("%v", maxSubArray(input))
}

func maxSubArray(nums []int) int {
	start := 0
	end := 1
	sum := nums[start]
	max := sum
	for end < len(nums) {
		if nums[end] > sum+nums[end] {
			sum = nums[end]
			start = end
		} else {
			sum += nums[end]
		}
		fmt.Printf("%v \n", sum)
		if sum > max {
			max = sum
		}
		end++
	}
	return max
}
