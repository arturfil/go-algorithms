package workingproblem

import (
	"fmt"
	"sort"
)

func Merge(intervals [][]int) [][]int {
	sort.Slice(intervals, func(a, b int) bool {
		return intervals[a][0] < intervals[b][0]
	})
	res := [][]int{intervals[0]}

	for _, inter := range intervals[1:] {
		last := res[len(res)-1]
		if inter[0] <= last[1] {
			maxV := max(inter[1], last[1])
			fmt.Println("maxV", maxV)
			last[1] = maxV
		} else {
			res = append(res, inter)
		}
	}

	fmt.Println("res ->", res)
	return res
}
