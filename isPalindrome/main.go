package main

import (
	"fmt"
	"strconv"
)

func main() {
	//var num = 1000021
	var num = 12321
	expected := true
	got := isPalindrome(num)

	if got == expected {
		fmt.Printf("number: %v is a palidrome\n", num)
	} else {
		fmt.Printf("number: %v is not a palidrome\n", num)
	}
}

func isPalindrome(num int) bool {

	// convert int to slice of runes
	runes := []rune(strconv.Itoa(num))

	// reverse the runes in the slice
	for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
		fmt.Printf("runes[i]: %v\n", runes[i])
		fmt.Printf("runes[j]: %v\n", runes[j])
		runes[i], runes[j] = runes[j], runes[i]
		fmt.Println("assignment-----------")
		fmt.Printf("runes[i]: %v\n", runes[i])
		fmt.Printf("runes[j]: %v\n", runes[j])
		fmt.Println("end------------------")
	}

	// convert back to int, via string
	y, err := strconv.Atoi(string(runes))

	// check for an error (overflow) and equality
	if err == nil && num == y {
		return true
	}

	return false
}

func isPalindrome1(num int) bool {
	if num < 0 {
		return false
	}

	digitCount, digitCounter := 0, num
	for digitCounter != 0 {
		digitCounter = digitCounter / 10
		digitCount++
	}

	var iterations = digitCount / 2

	var leftModulo = 1
	var rightModulo = 10

	for i := 0; i < digitCount; i++ {
		leftModulo *= 10
	}

	leftDigit := (num % leftModulo) / (leftModulo / 10)
	rightDigit := num % rightModulo

	for i := 0; i < iterations; i++ {
		if leftDigit != rightDigit {
			return false
		}

		leftModulo = leftModulo / 10
		rightModulo = rightModulo * 10

		leftDigit = (num % leftModulo) / (leftModulo / 10)
		rightDigit = (num % rightModulo) / (rightModulo / 10)
	}

	return true
}
