package main

import (
	"fmt"
	"strings"
)

func numUniqueEmails(emails []string) int {
	mapUnique := make(map[string]bool)
	for i, s := range emails {
		arr1 := strings.Split(s, "@")
		s1, dom := arr1[0], arr1[1]
		s1 = strings.Split(s1, "+")[0]
		s1 = strings.ReplaceAll(s1, ".", "")
		s1 = strings.Join([]string{s1, dom}, "@")
		mapUnique[s1] = true
		fmt.Println(s1, i)
	}
	return len(mapUnique)
}
