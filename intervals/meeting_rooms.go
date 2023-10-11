package intervals

import "sort"

func CanAttendMeetings(intervals [][]int) bool {
    // solution, iterate through the meetings and if the end of one is 
    // larger than the beg of one, return false, else true
    sort.Ints(intervals[0])
    
    for i := 0; i < len(intervals)-2; i++ {
        if intervals[i][1] < intervals[i+1][0] {
            return false
        }
    }
    
    return true
}
