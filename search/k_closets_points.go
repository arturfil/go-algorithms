package search

import (
	"fmt"
	"sort"
)

type Value struct {
    distance int
    point []int
}

func KClosest(points [][]int, k int) [][]int {
    distance, res := []Value{}, [][]int{}

    for _, point := range points {
        dist := (point[0] * point[0]) + (point[1] * point[1])
        fmt.Println(dist)
        distance = append(distance, Value{dist, point})
    }

    sort.Slice(distance, func(a, b int) bool {
        return distance[a].distance < distance[b].distance
    })

    for i := 0; i < k; i++ {
        res = append(res, distance[i].point)
    } 

    return res
}
