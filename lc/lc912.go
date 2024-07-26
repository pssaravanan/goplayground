package lc

import "fmt"

func SortArray(nums []int) []int {
	if len(nums) < 2 {
		return nums
	} else if len(nums) == 2 {
		if nums[0] > nums[1] {
			return []int{nums[1], nums[0]}
		} else {
			return []int{nums[0], nums[1]}
		}
	} else {
		mid := len(nums) / 2
		leftarray := SortArray(nums[:mid])
		rightarray := SortArray(nums[mid:])
		var result []int
		i := 0
		j := 0
		fmt.Println(nums, leftarray, rightarray)
		for i < len(leftarray) || j < len(rightarray) {
			fmt.Println("index", i, j)
			if j >= len(rightarray) || (i < len(leftarray) && leftarray[i] < rightarray[j]) {
				result = append(result, leftarray[i])
				i = i + 1
			} else if i >= len(leftarray) || (j < len(rightarray) && rightarray[j] < leftarray[i]) {
				result = append(result, rightarray[j])
				j = j + 1
			}
		}
		fmt.Println("index", i, j)
		return result
	}
}
