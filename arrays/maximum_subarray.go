package arrays

func maxSubArray(nums []int) int {
    maxV, maxTemp := nums[0], nums[0]

    for _, num := range nums {
        maxTemp = max(num, maxTemp + num)
        maxV = max(maxV, maxTemp)
    }

    return maxV
}

func max(vals ...int) int {
    maxv := vals[0]
    for _, val := range vals {
        if val > maxv {
            maxv = val
        }
    }
    return maxv
}


