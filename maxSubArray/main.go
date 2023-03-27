package main

import "fmt"

func main() {
	nums := []int{-2, 1, -3, 4, -1, 2, 1, -5, 4}
	expected := 6

	got := maxSubArray1(nums)

	if got == expected {
		fmt.Println("correct answer: ", got, expected)
	} else {
		fmt.Println("wrong answer: ", got, expected)
	}
}

func maxSubArray1(nums []int) int {
	return findMaxSubArray(nums, 0, len(nums)-1)
}

func findMaxSubArray(nums []int, left, right int) int {
	if left == right {
		return nums[left]
	}

	mid := (left + right) / 2

	//return max(findMaxSubArray(nums, left, mid), maxCrossing(nums, left, right, mid), findMaxSubArray(nums, mid+1, right))
	leftMax := findMaxSubArray(nums, left, mid)
	rightMax := findMaxSubArray(nums, mid+1, right)
	crossMax := maxCrossing(nums, left, right, mid)
	return max(leftMax, rightMax, crossMax)
}

func maxCrossing(nums []int, left, right, mid int) int {
	leftSum, rightSum, sum := -1<<31, -1<<31, 0

	for i := mid; i >= left; i-- {
		sum += nums[i]
		leftSum = max(leftSum, sum)
	}
	sum = 0
	for i := mid + 1; i < right; i++ {
		sum += nums[i]
		rightSum = max(rightSum, sum)
	}

	return leftSum + rightSum
}

func max(values ...int) int {
	var maxval int
	for i := 0; i < len(values); i++ {
		if values[i] > maxval {
			maxval = values[i]
		}
	}
	return maxval
}

func maxSubArray(nums []int) int {
	if len(nums) == 0 {
		return 0
	}
	if len(nums) == 1 {
		return nums[0]
	}
	dp, res := make([]int, len(nums)), nums[0]
	dp[0] = nums[0]
	for i := 1; i < len(nums); i++ {
		if dp[i-1] > 0 {
			dp[i] = nums[i] + dp[i-1]
		} else {
			dp[i] = nums[i]
		}
		res = max(res, dp[i])
	}

	return res
}
