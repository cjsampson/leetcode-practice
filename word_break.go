package main

import "fmt"

func main() {
	// var s = "applepenapple"
	// var wordDict = []string{"apple", "pen"}
	var s = "cars"
	var wordDict = []string{"cats", "dog", "sand", "and", "cat"}

	result := wordBreak(s, wordDict)
	if result == true {
		fmt.Println("incorrect result")
	} else {
		fmt.Println("correct result")
	}
}

func wordBreak(s string, wordDict []string) bool {
	return false
}
