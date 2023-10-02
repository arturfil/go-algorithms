package main

import (
	"fmt"
	"leetcode/intervals"
)

func main() {
    // helpers.PrintDoneProblems()
    inters := [][]int {
        {1,3},{2,6},{8,10},{15,18},
    }

    fmt.Println(intervals.Merge(inters))
}
