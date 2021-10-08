package main

import (
	"fmt"
	"strconv"
)

func main() {
	fmt.Println(isPalindrome(90))
	fmt.Println(isPalindrome(9009))
	fmt.Println(isPalindrome(-90))
	fmt.Println(isPalindrome(-12021))

}

func isPalindrome(x int) bool {
	str := strconv.Itoa(x)

	runes := []rune(str)
	for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
		runes[i], runes[j] = runes[j], runes[i]
	}

	rstr := string(runes)

	if str == rstr {
		return true
	} else {
		return false
	}

}
