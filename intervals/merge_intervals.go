package intervals

import "sort"

func Merge(intervals [][]int) [][]int {
    sort.Slice(intervals, func (a, b int) bool {
       return intervals[a][0] < intervals[b][0] 
    })
    res := [][]int {intervals[0]}

	for _, interval := range intervals[1:] {
        last := res[len(res)-1]
        if interval[0] <= last[1] {
            last[1] = max(interval[1], last[1])
        } else  {
            res = append(res, interval)
        }
    }

    return res
}

