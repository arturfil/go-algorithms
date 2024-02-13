package intervals

import (
	"sort"
)

func EraseOverlapIntervals(intervals [][]int) int {

    sort.Slice(intervals, func(a, b int) bool {
        return intervals[a][1] < intervals[b][1]        
    })

    counter, prev  := 0 ,intervals[0]

    for _, interval := range intervals[1:] {

        if  interval[0] >= prev[1] {
            prev = interval
        } else {
            counter++
        }
    }

    return counter
}

