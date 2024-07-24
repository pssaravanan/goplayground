package lc

import (
	"fmt"
	"sort"
	"strconv"
)

func SortJumbled(mapping []int, nums []int) []int {
	type mapV struct {
		num   int
		value int
	}
	var mapValues []mapV
	var map_a []int
	for i := range nums {
		fmt.Println(i)
		s := strconv.Itoa(nums[i])
		var map_s string
		fmt.Println(map_s)
		for j := range s {
			m_i, _ := strconv.Atoi(string(s[j]))
			fmt.Println(m_i)
			m_n := strconv.Itoa(mapping[m_i])
			map_s = map_s + m_n
		}
		fmt.Println("res", map_s)
		num, _ := strconv.Atoi(map_s)
		map_a = append(map_a, num)
	}
	for i := range nums {
		mapValues = append(mapValues, mapV{num: nums[i], value: map_a[i]})
	}
	sort.Slice(mapValues, func(ii, jj int) bool {
		return mapValues[ii].value < mapValues[jj].value
	})

	var result []int
	for i := range mapValues {
		result = append(result, mapValues[i].num)
	}
	return result
}
