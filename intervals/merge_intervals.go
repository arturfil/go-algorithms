package intervals

func Merge(intervals [][]int) [][]int {
    merged := [][]int {intervals[0]}

    for _, interval := range intervals {
        last := len(merged)-1 
        if interval[0] > merged[last][1] {
            merged = append(merged, interval)
        } else  {
            merged[last][1] = max(interval[1], merged[last][1])
        }
    }

    return merged
}

