package main

import "fmt"

// import "strings"

func main() {
	var paragraph = "Bob hit a ball, the hit BALL flew far after it was hit."
	var banned = []string{"hit"}
	var expected = "ball"

	fmt.Printf("banned: %v\n", banned)
	fmt.Printf("expected: %v\n", expected)

	for i := 0; i < len(paragraph); i++ {
		fmt.Printf("paragraph[i]: %v\n", paragraph[i])
	}
}
