package intervals

import (
	"sort"
)

func minMeetingRooms(intervals [][]int) int {
    size := len(intervals)
    starts, ends := make([]int, size), make([]int, size)

    for i := 0; i < size; i++ {
        starts[i] = intervals[i][0]
        ends[i] = intervals[i][1]
    }
    sort.Ints(starts)
    sort.Ints(ends)

    rooms, end := 0, 0
    for i := 0; i < len(intervals); i++ {
        if starts[i] < ends[end] {
            rooms++
        } else {
            end++
        }
    }

    return rooms
}
