package main

import (
	"fmt"
	"sort"
)

type byLength []string

func (s byLength) Len() int {
	return len(s)
}
func (s byLength) Swap(i, j int) {
	s[i], s[j] = s[j], s[i]
}
func (s byLength) Less(i, j int) bool {
	if len(s[i]) < len(s[j]) {
		return true
	} else if len(s[i]) == len(s[j]) {
		return s[i] < s[j]
	} else {
		return false
	}
}

func main() {
	fruits := []string{"apple", "peach", "banana", "kiwi", "apple"}
	sort.Sort(byLength(fruits))
	fmt.Println(fruits)
}
