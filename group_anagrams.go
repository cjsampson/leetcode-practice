package main

import "fmt"
import "strings"
import "sort"

func main() {
	input := []string{"eat", "tea", "tan", "ate", "nat", "bat"}
	// expected := [][]string{{"bat"}, {"nat", "tan"}, {"ate", "eat", "tea"}}

	fmt.Printf("input: %v\n", input)
	fmt.Println("--------")

	var result [][]string
	for i := 0; i < len(input); i++ {
		if len(result) == 0 {
			result = append(result, []string{input[i]})
			continue
		}

		for j := 0; j < len(result); j++ {
			if checkAnagram(result[j][0], input[i]) {
				result[j] = append(result[j], input[i])
				break
			}
			if j == len(result)-1 {
				result = append(result, []string{input[i]})
				break
			}
		}
	}
	fmt.Printf("result: %v\n", result)
}

func checkAnagram(str1, str2 string) bool {
	if len(str1) != len(str2) {
		return false
	}
	split1 := strings.Split(str1, "")
	split2 := strings.Split(str2, "")

	sort.Strings(split1)
	sort.Strings(split2)

	join1 := strings.Join(split1, "")
	join2 := strings.Join(split2, "")

	if join1 == join2 {
		return true
	}

	return false
}
