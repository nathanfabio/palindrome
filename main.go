package main

import (
	"fmt"
	"strings"
)

func main() {

	fmt.Printf("Is the word 'kayak' a palindrome? %v\n", isItAPalindrome("kayak"))
	fmt.Printf("Is the word 'palindrome' a palindrome? %v\n", isItAPalindrome("palindrome"))
}

func isItAPalindrome(word string) bool {
	clearSpace := strings.ReplaceAll(word, " ", "")
	lowerCase:= strings.ToLower(clearSpace)
	
	left:= 0
	right:= len(lowerCase) - 1
	
	for left < right {
		if lowerCase[left] != lowerCase[right] {
			return false
		}
		left++
		right--
	}
	return true
}