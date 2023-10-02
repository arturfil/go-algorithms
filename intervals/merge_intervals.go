package intervals

import "sort"

func Merge(intervals [][]int) [][]int {
    sort.Slice(intervals, func (a, b int) bool {
       return intervals[a][0] < intervals[b][0] 
    })
    res := [][]int {intervals[0]}

    for _, interval := range intervals {
        last := len(res)-1 
        if interval[0] > res[last][1] {
            res = append(res, interval)
        } else  {
            res[last][1] = max(interval[1], res[last][1])
        }
    }

    return res
}

