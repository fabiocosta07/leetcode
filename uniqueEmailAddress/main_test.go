package main

import (
	"testing"
)

func TestUnique1(t *testing.T) {
	input := []string{"test.email+alex@leetcode.com", "test.e.mail+bob.cathy@leetcode.com", "testemail+david@lee.tcode.com"}

	r := numUniqueEmails(input)
	if r != 2 {
		t.Errorf("Error")
	}
}

func TestUnique2(t *testing.T) {
	input := []string{"a@leetcode.com", "b@leetcode.com", "c@leetcode.com"}
	r := numUniqueEmails(input)
	if r != 3 {
		t.Errorf("Error")
	}
}

func TestUnique3(t *testing.T) {
	input := []string{"test.email+alex@leetcode.com", "test.email@leetcode.com"}
	r := numUniqueEmails(input)
	if r != 1 {
		t.Errorf("Error")
	}
}
