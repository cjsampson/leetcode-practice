package main

import (
	"fmt"
	"math"
	"strings"
)

func main() {
	strs := []string{"flower", "flow", "flight"}
	//strs := []string{"dog", "racecar", "car"}

	longestCommonPrefix(strs)
}

func longestCommonPrefix(strs []string) string {
	var prefix strings.Builder
	if len(strs) == 0 {
		return prefix.String()
	}
	var lp int
	fmt.Println("---------------")
	fmt.Printf("lp:%v\n", lp)
	fmt.Println("---------------")
	for i := 1; i < len(strs)-1; i++ {
		for j := 0; j < int(
			math.Min(
				float64(len(strs[i])), float64(len(strs[i+1])),
			)); j++ {

			if strs[i][j] != strs[i+1][j+1] {
				continue
			}

			//if lp <= prefix.WriteByte(strs[i][j]) {
			//
			//}

			fmt.Printf("string(strs[i][j]): %v\n", string(strs[i][j]))
			fmt.Printf("string(strs[i+1][j+1]): %v\n", string(strs[i+1][j+1]))
			//if strs[i][j] == strs[i+1][j+1] {
			//	prefix.WriteByte(strs[i][j])
			//}
		}
	}

	return prefix.String()
}
