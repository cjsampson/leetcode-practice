package main

import "fmt"

func main() {
	haystack := "hello"
	needle := "ll"
	result := strStr(haystack, needle)

	if result == 2 {
		fmt.Printf("correct result: %v\n", result)
	} else {
		fmt.Printf("incorrect result: %v\n", result)
	}
}

func strStr(haystack string, needle string) int {

	return 0
}
