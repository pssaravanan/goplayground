package lc

import (
	"sort"
)

type nameHeight struct {
	name   string
	height int
}

func SortPeople(names []string, heights []int) []string {

	var nameHeights []nameHeight
	for i := range names {
		nameHeights = append(nameHeights, nameHeight{name: names[i], height: heights[i]})
	}
	sort.Slice(nameHeights, func(i, j int) bool {
		return nameHeights[i].height > nameHeights[j].height
	})
	var result []string
	for i := range nameHeights {
		result = append(result, nameHeights[i].name)
	}
	return result
}
