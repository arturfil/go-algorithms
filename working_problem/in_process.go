package workingproblem

import "fmt"

func ProductExceptSelf(nums []int) []int {
    res := make([]int, len(nums))

    prev := 1 
    for i, num := range nums {
        res[i] = prev
        prev *= num
    }
    
    prev = 1 
    for i := len(nums)-1; i >= 0; i-- {
        res[i] *= prev
        prev *= nums[i]
    }

    fmt.Println("res->", res)
	return res
}
