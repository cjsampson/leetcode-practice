package main

import "fmt"

func main() {
	list := []int{1, 2, 3, 1}
	result := containsDuplicate(list)
	fmt.Println("---------------")
	fmt.Printf("result:%v\n", result)
	fmt.Println("---------------")
}

func containsDuplicate(nums []int) bool {
	var list []int
	for i := 0; i < len(nums); i++ {
		if list[nums[i]] == 0 {
			list[nums[i]]++
		} else {
			return true
		}
	}

	return false
}
