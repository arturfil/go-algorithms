package dynamicprogramming

import (
	"math"
)

func coinChange(coins []int, amount int) int {
   dp := make([]int, amount + 1)
   for i := range dp { dp[i] = math.MaxInt32 }
   dp[0] = 0

   for _, coin := range coins {
       for i := coin; i <= amount; i++ {
           dp[i] = Min(dp[i], dp[i - coin] + 1) // 1 signifies the coin itself
       }
   }

   if dp[amount] == math.MaxInt32 {
       return -1
   } else {
       return dp[amount]
   }
}

func Min(a, b int) int {
    if a < b {
        return a
    } else {
        return b
    }
}
