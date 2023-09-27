package main

import (
	"fmt"
	"leetcode/search"
)


func main() {
    // helpers.RandomProblem()
    fmt.Println(search.KClosest([][]int{{1,2}, {-1,2}, {3,5}, {3,6}, {2, 4}}, 2))
}
