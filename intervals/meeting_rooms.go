package intervals

import "sort"

func CanAttendMeetings(intervals [][]int) bool {
    sort.Slice(intervals, func(a, b int) bool {
        return intervals[a][0] < intervals[b][0]
    }) 
    
    for i := 0; i < len(intervals)-1; i++ {
        if intervals[i][1] > intervals[i+1][0] {
            return false
        }
    }
    
    return true
}
