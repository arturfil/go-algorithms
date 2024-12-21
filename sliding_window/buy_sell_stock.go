package slidingwindow

import "math"

func MaxProfit(prices []int) int {
    l, r, profit := 0, 1, 0
    
    for r < len(prices) {

        if prices[l] < prices[r] {
            curr := prices[r] - prices[l]
            profit = Max(curr, profit)

        } else {
            l = r
        }
        r++
    }
    return profit
}

func maxProfit(prices []int) int {
    min, max := math.MaxInt, 0
    for _, price := range prices {
        if price < min {
            min = price
        } else {
            max = Max(max, price - min)
        }
    }

    return max
}


func Max(args ...int) int {
    maxVal := args[0]
    for _, n := range args {
        if n > maxVal { maxVal = n }
    }

    return maxVal
}
