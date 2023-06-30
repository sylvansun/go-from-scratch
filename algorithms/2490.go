package main

import (
	"fmt"
)

func isCircularSentence(sentence string) bool {
	length := len(sentence)
	if sentence[0] != sentence[length-1] {
		return false
	}
	for i := 1; i < length-1; i++ {
		if sentence[i] == ' ' && (sentence[i-1] != sentence[i+1]) {
			return false
		}
	}
	return true
}

func main() {
	fmt.Println(isCircularSentence("A man, a plan, a canal: Panama"))
}
