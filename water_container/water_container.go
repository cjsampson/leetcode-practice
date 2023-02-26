package main

import "fmt"
import "math"

func main() {
	// heights := []int{1, 8, 6, 2, 5, 4, 8, 3, 7}
	// expected := 49
	heights := []int{1, 1}
	expected := 1

	fmt.Println("---------")
	fmt.Printf("heights: %v\n\n", heights)
	fmt.Printf("expected: %v\n", expected)

	result := waterContainer(heights)

	if result != expected {
		fmt.Println("------incorrect--------")
		fmt.Printf("result: %v\n", result)
	} else {
		fmt.Println("------correct--------")
		fmt.Printf("result: %v\n", result)
	}
}

func waterContainer(heights []int) int {
	var height float64
	for i := 0; i < len(heights); i++ {
		if i == len(heights)-2 {
			break
		}
		for j := i + 1; i < len(heights); j++ {
			distanceBetween := j - i
			fmt.Printf("distanceBetween: %v\n", distanceBetween)
			minHeight := math.Min(float64(heights[i]), float64(heights[j]))
			fmt.Printf("minHeight: %v\n", minHeight)
			height = math.Max(float64(height), float64(minHeight*float64(distanceBetween)))
			fmt.Printf("height: %v\n", height)

			if j == len(heights)-1 {
				break
			}
		}
	}

	return int(height)
}
