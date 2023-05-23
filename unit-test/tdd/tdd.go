package tdd

import "strings"

var vowels = map[string]bool{
	"A": true,
	"E": true,
	"I": true,
	"O": true,
	"U": true,
}

func VowelCount(sentence string) uint {
	var count uint
	sentence = strings.ToUpper(sentence)
	for _, char := range sentence {
		if vowels[string(char)] {
			count++
		}
	}
	return count
}
