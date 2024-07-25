package main

import (
	"fmt"

	"github.com/pssaravanan/goplayground/lc"
)

func main() {
	// result := lc.SortPeople([]string{"Mary", "John", "Emma"}, []int{180, 165, 170})
	// fmt.Println(result)
	// lc.ReverseParentheses("(ed(et(oc))el)")
	// lc.SortJumbled([]int{8,9,4,0,2,1,3,5,7,6}, []int{991,338,38})
	result := lc.SortArray([]int{5, 1, 1, 2, 0, 0})
	fmt.Println(result)
}
