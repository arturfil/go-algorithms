package main

import (
	"fmt"
	"leetcode/graphs"
)


func main() {
    // helpers.RandomProblem()
    res := graphs.CanFinishBFS(11, [][]int{{1,2}, {2,3}, {3,4}, {4,0}})

    fmt.Println(res)
}
