package greedy

func MaxSubArray(nums []int) int {

    tot, max := nums[0], nums[0] 
    for _, num := range nums[1:] {
        max = Max(num, max + num)
        tot = Max(tot, max)
    }

    return tot
}

func Max(nums ...int) int {
    max := nums[0]
    for _, num := range nums {
        if num > max { max = num }
    }

    return max
}
