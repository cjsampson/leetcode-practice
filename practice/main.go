package main

import "fmt"

func main() {
	//str1 := "there is a thing over there    "
	str := "Hello World  "
	//str := "a"
	length := lengthOfLastWord(str)
	fmt.Println("---------------")
	fmt.Printf("length:%v\n", length)
	fmt.Println("---------------")
}

func lengthOfLastWord(str string) int {
	n := len(str) - 1
	for n >= 0 && str[n] == ' ' {
		n--
	}

	var length int
	for n >= 0 && str[n] != ' ' {
		n--
		length++
	}

	return length
}
