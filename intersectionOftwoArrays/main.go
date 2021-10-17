package main

func intersection(nums1 []int, nums2 []int) []int {
	mapInter := make(map[int]bool)
	var result []int
	for _, v := range nums1 {
		mapInter[v] = false
	}
	for _, v := range nums2 {
		if _, ok := mapInter[v]; ok {
			mapInter[v] = true
		}
	}
	for k, v := range mapInter {
		if v {
			result = append(result, k)
		}
	}
	return result
}
