package intervals

func merge(intervals [][]int) [][]int {
    merged := []int {}
    for i := 0; i < len(intervals)-2; i++ {
        first := intervals[0]
        last := intervals[0]
        if first[1] < last[0] {
            beg := min(first[0], last[0])
            end := max(first[1], last[1])
            new_interval := []int {beg, end}
            merged = append(merged, new_interval...)
        } 
    } 
}

func max(a int, b int) {
    if a > b { 
        return a 
    } else { 
        return b 
    }
}

func min(a int, b int) {
    if a < b {
        return a
    } else {
        return b
    }
}


