package main

func firstUniqChar(s string) int {
	charArr := []rune(s)

	mapUnique := make(map[rune]bool)
	for i := range charArr {
		char := charArr[i]
		if _, ok := mapUnique[char]; ok {
			mapUnique[char] = true
		} else {
			mapUnique[char] = false
		}
	}
	min := -1
	for i := range charArr {
		char := charArr[i]
		if !mapUnique[char] {
			if min < 0 {
				min = i
			} else if i < min {
				min = i
			}
		}
	}
	return min
}
