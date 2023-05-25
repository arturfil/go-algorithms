package binary_search

import "fmt"

// example 1
// [4,5,6,7,0,1]

/*
   goal? -> to find the min val
   how? -> when the left value is larger than the right
   possible impl -> use binary search and compare left to right value
*/

func FindMin(nums []int) int {
    left, right := 0, len(nums)-1
    for left < right {
        mid := (right + left) / 2
        if nums[mid] > nums[right] {
           left = mid + 1
        } else {
            right = mid
        }
    }
    return nums[left] 
}
