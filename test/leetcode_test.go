package test

import (
	"fmt"
	"sort"
	"testing"
	"unicode"
)

func TestLC(t *testing.T) {
	s1 := "X2Y3XZ"
	fmt.Printf("input: %s, output: %s\n", s1, countStr(s1))
	s2 := "Z3X(XY)2"
	fmt.Printf("input: %s, output: %s\n", s2, countStr(s2))
	s3 := "Z4(Y2(XZ2)3)2X2"
	fmt.Printf("input: %s, output: %s\n", s3, countStr(s3))
}

func countLetters(s string) map[rune]int {
	counts := make(map[rune]int)
	for _, r := range s {
		if unicode.IsLetter(r) {
			counts[unicode.ToUpper(r)]++
		}
	}
	return counts
}

func countStr(s string) string {
	var stack []string
	var result string

	for i := 0; i < len(s); i++ {
		if s[i] == '(' {
			stack = append(stack, "")
			continue
		} else if s[i] == ')' {
			sub := stack[len(stack)-1]
			stack = stack[:len(stack)-1]
			num := 0
			for i+1 < len(s) && unicode.IsDigit(rune(s[i+1])) {
				num = num*10 + int(s[i+1]-'0')
				i++
			}
			counts := countLetters(sub)
			var keys []rune
			for k := range counts {
				keys = append(keys, k)
			}
			sort.Slice(keys, func(i, j int) bool {
				return keys[i] < keys[j]
			})
			for _, k := range keys {
				result += fmt.Sprintf("%c%d", k, counts[k]*num)
			}
			continue
		}

		if len(stack) > 0 {
			stack[len(stack)-1] += string(s[i])
		} else {
			result += string(s[i])
		}
	}

	counts := countLetters(result)
	var keys []rune
	for k := range counts {
		keys = append(keys, k)
	}
	sort.Slice(keys, func(i, j int) bool {
		return keys[i] < keys[j]
	})
	result = ""
	for _, k := range keys {
		result += fmt.Sprintf("%c%d", k, counts[k])
	}
	return result
}
