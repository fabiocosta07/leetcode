package main

func main() {

}

func topKFrequent(nums []int, k int) []int {

	counterMap := make(map[int]int)

	for _, v := range nums {
		counterMap[v] += 1
	}

	freqArr := make([][]int, len(nums)+1)

	for k, v := range counterMap {
		freqArr[v] = append(freqArr[v], k)
	}

	res := make([]int, 0)
	for i := len(freqArr) - 1; i >= 0; i-- {
		res = append(res, freqArr[i]...)
	}

	return res[:k]
}
