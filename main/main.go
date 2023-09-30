package main

import (
	"fmt"
	"leetcode/intervals"
)

func main() {
    // helpers.PrintDoneProblems()
    inters := [][]int {
        {1,3},{2,4},{5,8},{9,10},{10,11},
    }

    fmt.Println(intervals.Merge(inters))
}
