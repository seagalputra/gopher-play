package kata

import (
	"strings"
)

func duplicate_count(s1 string) int {
	lower := strings.ToLower(s1)
	dictionary := map[string]int{}
	chars := strings.Split(lower, "")

	for _, char := range chars {
		if _, isExist := dictionary[char]; isExist {
			dictionary[char] = dictionary[char] + 1
		} else {
			dictionary[char] = 1
		}
	}

	counter := 0
	for _, value := range dictionary {
		if value > 1 {
			counter++
		}
	}

	return counter
}
