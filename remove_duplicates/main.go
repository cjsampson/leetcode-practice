package main

import "fmt"

func main() {
	duplicates := []int{1, 1, 2, 2, 3, 3, 4}

	removeDuplicates(duplicates)

	fmt.Printf("duplicates: %v\n", duplicates)
}

func removeDuplicates(nums []int) int {
	var i int

	for j := 1; j < len(nums); j++ {
		if nums[j] != nums[i] {
			i++
			nums[i] = nums[j]
			fmt.Printf("i: %v\n", i)
			fmt.Printf("j: %v\n", j)
			fmt.Printf("nums: %v\n", nums)
		}
	}

	return i + 1
}
