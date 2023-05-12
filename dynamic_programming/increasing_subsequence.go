package dynamicprogramming

import (
	"sort"
)

// O(n^2) answer
func LengthOfLIS(nums []int) int {
    dp := make([]int, len(nums))
    lis := 1 // len of increase subs

    for i := 0; i < len(nums); i++ {
        dp[i] = 1
        for j := 0; j < i; j++ {
            if nums[j] < nums[i] && 1 + dp[j] > dp[i] {
                dp[i] = dp[j] + 1
            }
        }
        if dp[i] > lis { lis = dp[i] }
    }
    return lis
}

// O(n log(n) ) answer
func LengthOfLISAlt(nums []int) int {
    dp := []int{}

    for _, num := range nums {
        i := sort.SearchInts(dp, num)
        if i == len(dp) {
            dp = append(dp, num)
        } else {
            dp[i] = num
        }
    }
    return len(dp)
}
