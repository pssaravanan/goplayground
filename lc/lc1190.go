package lc

import "fmt"

// import "fmt"

func reverseString(s string) string {
	runes := []rune(s)
	for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
		runes[i], runes[j] = runes[j], runes[i]
	}
	return string(runes)
}

func ReverseParentheses(s string) string {
	fmt.Println(s)
	type rng struct {
		i int
		j int
	}
	var rngarray []rng
	var indexstack []int
	for i := range s {
		if s[i] == '(' {
			indexstack = append(indexstack, i)
		} else if s[i] == ')' {
			indexstacksize := len(indexstack)
			rngarray = append(rngarray, rng{i: indexstack[indexstacksize-1], j: i})
			indexstack = indexstack[:indexstacksize-1]
		}
		// fmt.Println(rngarray)
	}
	for i := range rngarray {
		ss := s[rngarray[i].i+1 : rngarray[i].j]
		s = s[:rngarray[i].i+1] + reverseString(ss) + s[rngarray[i].j:]
	}
	var result string
	for i := range s {
		if s[i] != '(' && s[i] != ')' {
			result = result + string(s[i])
		}
	}
	return result
}
