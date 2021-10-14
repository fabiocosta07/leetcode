package main

import "fmt"

const s1 = "()"
const s2 = "()[]{}"
const s3 = "(]"
const s4 = "([)]"
const s5 = "{[]}"
const s6 = "]"

func isValid(s string) bool {
	var stack []rune
	for _, c := range s {
		if c == '{' || c == '(' || c == '[' {
			stack = append(stack, c)
		} else if c == '}' || c == ')' || c == ']' {
			if len(stack) == 0 {
				return false
			}
			switch stack[len(stack)-1] {
			case '{':
				if c != '}' {
					return false
				}
			case '(':
				if c != ')' {
					return false
				}
			case '[':
				if c != ']' {
					return false
				}
			}
			stack = stack[:len(stack)-1]
		}
	}
	return len(stack) == 0
}
func main() {
	fmt.Printf("s1: %t ", isValid(s1))
	fmt.Printf("s2: %t ", isValid(s2))
	fmt.Printf("s3: %t ", isValid(s3))
	fmt.Printf("s4: %t ", isValid(s4))
	fmt.Printf("s5: %t ", isValid(s5))
	fmt.Printf("s6: %t ", isValid(s6))
}
